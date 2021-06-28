package handlers

import (
	"eCommerce/database"
	"encoding/json"
	"net/http"
)

func GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(database.GetProductList())
	}
}
