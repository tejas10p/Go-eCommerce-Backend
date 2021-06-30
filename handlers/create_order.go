package handlers

import (
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"net/http"
)

func CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var newOrder models.Order
		_ = json.NewDecoder(r.Body).Decode(&newOrder)
		json.NewEncoder(w).Encode(database.AddOrder(newOrder))
	}
}
