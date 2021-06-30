package handlers

import (
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"net/http"
)

func AddProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var newProduct models.Product
		_ = json.NewDecoder(r.Body).Decode(&newProduct)
		json.NewEncoder(w).Encode(database.AddProduct(newProduct))
	}
}
