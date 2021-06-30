package handlers

import (
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"net/http"
)

func AddItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var newItem models.Item
		_ = json.NewDecoder(r.Body).Decode(&newItem)
		json.NewEncoder(w).Encode(database.AddItem(newItem))
	}
}
