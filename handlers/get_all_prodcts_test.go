package handlers

import (
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestNilForNoProducts(t *testing.T) {
	database.Init()
	r := httptest.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	GetProducts()(w, r)
	var result []models.Product
	_ = json.NewDecoder(w.Body).Decode(&result)
	assert.Equal(t, nil, result)
	database.Close()
}
