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
		newProduct.ID = database.GetId()
		if !newProduct.Validate() {
			json.NewEncoder(w).Encode(models.Product{})
			return
		}
		database.Products = append(database.Products, newProduct)
		json.NewEncoder(w).Encode(newProduct)
	}
}
