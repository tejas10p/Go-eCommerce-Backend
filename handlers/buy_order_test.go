package handlers

import (
	"database/sql"
	"eCommerce/database"
	"eCommerce/kafka"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CreateBuyOrderTest struct {
	suite.Suite

	dbx *sql.DB
}

func (suite *CreateBuyOrderTest) SetupSuite() {
	database.Init()
	suite.dbx = database.Db
	SetDummyEntries(suite.dbx)
	kafka.CreateTopic("producers", 1)
}

func (suite *CreateBuyOrderTest) TearDownSuite() {
	ClearDb(suite.dbx)
	suite.dbx.Close()
}

func TestBuyOrderSuite(t *testing.T) {
	suite.Run(t, new(CreateBuyOrderTest))
}

func (suite *CreateBuyOrderTest) TestStatusCheckFailure() {
	t := suite.T()

	r := httptest.NewRequest("PUT", "/orders/buy/2", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "2"})
	w := httptest.NewRecorder()

	BuyOrder()(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func (suite *CreateBuyOrderTest) TestOrderPurchaseSuccess() {
	t := suite.T()

	r := httptest.NewRequest("PUT", "/orders/buy/1", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	BuyOrder()(w, r)

	var result int
	_ = json.NewDecoder(w.Body).Decode(&result)

	assert.Equal(t, http.StatusAccepted, w.Code)
	assert.Equal(t, 9500, result)
}

func SetDummyEntries(db *sql.DB) {
	_, err := db.Exec("INSERT INTO user VALUES (?, ?, ?, ?)", "Tejas Pandey", "tejaspandey10@gmail.com", 99999999, "Somewhere, World")
	if err != nil {
		log.Fatalf("Could not fill entries - %s", err.Error())
	}

	_, err = db.Exec("INSERT INTO product (name, price, category) VALUES (?, ?, ?)", "Rice", 100, "Food")
	if err != nil {
		log.Fatalf("Could not fill entries - %s", err.Error())
	}

	_, err = db.Exec("INSERT INTO orders (email, address, status) VALUES (?, ?, ?)", "tejaspandey10@gmail.com", "Somewhere, World", "Created")
	if err != nil {
		log.Fatalf("Could not fill entries - %s", err.Error())
	}

	_, err = db.Exec("INSERT INTO item VALUES (?, ?, ?, ?)", 1, 1, 100, 9000)
	if err != nil {
		log.Fatalf("Could not fill entries - %s", err.Error())
	}

	_, err = db.Exec("INSERT INTO item VALUES (?, ?, ?, ?)", 1, 1, 5, 500)
	if err != nil {
		log.Fatalf("Could not fill entries - %s", err.Error())
	}
}

func ClearDb(db *sql.DB) {
	_, err := db.Exec("DELETE FROM item")
	if err != nil {
		log.Fatalf("Error clearing dummy entries - %s", err.Error())
	}

	_, err = db.Exec("DELETE FROM product")
	if err != nil {
		log.Fatalf("Error clearing dummy entries - %s", err.Error())
	}

	_, err = db.Exec("DELETE FROM orders")
	if err != nil {
		log.Fatalf("Error clearing dummy entries - %s", err.Error())
	}

	_, err = db.Exec("DELETE FROM user")
	if err != nil {
		log.Fatalf("Error clearing dummy entries - %s", err.Error())
	}
}
