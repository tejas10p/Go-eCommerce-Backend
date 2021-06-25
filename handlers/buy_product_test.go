package handlers

import (
	"bytes"
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func jsonReader(data interface{}) io.Reader {
	marshalledBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return bytes.NewReader(marshalledBytes)
}

func SetupTest() {
	user := models.User{FirstName: "Tejas", LastName: "Pandey", Email: "tejas@pan.com", Address: "placeholder"}
	database.Users = append(database.Users, user)
	firstProduct := models.Product{ID: "3", Name: "RICE", Price: 150, Quantity: 250, Category: "FOOD"}
	database.Products = append(database.Products, firstProduct)
	secondProduct := models.Product{ID: "1", Name: "NOODLES", Price: 50, Quantity: 50, Category: "FOOD"}
	database.Products = append(database.Products, secondProduct)
}

func TestCustomerValidationFalse(t *testing.T) {
	SetupTest()
	user := models.User{FirstName: "Tejas", LastName: "Pandey", Email: "tej@pan.com", Address: "placeholder"}
	body := models.Order{ProductId: []string{"1", "2"}, Quantity: []int{2, 4}, Customer: &user}
	r := httptest.NewRequest("PUT", "/products/buy", jsonReader(body))
	w := httptest.NewRecorder()
	BuyProduct()(w, r)
	var isCustomerFound string
	_ = json.NewDecoder(w.Body).Decode(&isCustomerFound)
	assert.Equal(t, "Customer not found", isCustomerFound)
}

func TestCustomerValidationTrue(t *testing.T) {
	user := models.User{FirstName: "Tejas", LastName: "Pandey", Email: "tejas@pan.com", Address: "placeholder"}
	body := models.Order{ProductId: []string{"4", "2"}, Quantity: []int{2, 4}, Customer: &user}
	r := httptest.NewRequest("PUT", "/products/buy", jsonReader(body))
	w := httptest.NewRecorder()
	BuyProduct()(w, r)
	var modifiedProducts []models.Product
	_ = json.NewDecoder(w.Body).Decode(&modifiedProducts)
	assert.Equal(t, []models.Product(nil), modifiedProducts)
}

func TestQuantityIsLessThanRequired(t *testing.T) {
	user := models.User{FirstName: "Tejas", LastName: "Pandey", Email: "tejas@pan.com", Address: "placeholder"}
	body := models.Order{ProductId: []string{"1", "3"}, Quantity: []int{75, 251}, Customer: &user}
	r := httptest.NewRequest("PUT", "/products/buy", jsonReader(body))
	w := httptest.NewRecorder()
	BuyProduct()(w, r)
	var modifiedProducts []models.Product
	_ = json.NewDecoder(w.Body).Decode(&modifiedProducts)
	assert.Equal(t, []models.Product(nil), modifiedProducts)
}

func TestQuantityIsMoreThanOrEqualRequired(t *testing.T) {
	user := models.User{FirstName: "Tejas", LastName: "Pandey", Email: "tejas@pan.com", Address: "placeholder"}
	body := models.Order{ProductId: []string{"1", "3"}, Quantity: []int{10, 250}, Customer: &user}
	r := httptest.NewRequest("PUT", "/products/buy", jsonReader(body))
	w := httptest.NewRecorder()
	BuyProduct()(w, r)
	var requiredResult []models.Product
	firstProduct := models.Product{ID: "3", Name: "RICE", Price: 150, Quantity: 0, Category: "FOOD"}
	requiredResult = append(requiredResult, firstProduct)
	secondProduct := models.Product{ID: "1", Name: "NOODLES", Price: 50, Quantity: 40, Category: "FOOD"}
	requiredResult = append(requiredResult, secondProduct)
	var modifiedProducts []models.Product
	_ = json.NewDecoder(w.Body).Decode(&modifiedProducts)
	assert.Equal(t, requiredResult, modifiedProducts)
}