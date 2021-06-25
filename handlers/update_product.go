package handlers

import (
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for i, v := range database.Products {
			if v.ID == params["id"] {
				database.Products = append(database.Products[:i], database.Products[i+1:]...)
				break
			}
		}
		var newProduct models.Product
		_ = json.NewDecoder(r.Body).Decode(&newProduct)
		newProduct.ID = params["id"]
		database.Products = append(database.Products, newProduct)
		json.NewEncoder(w).Encode(newProduct)
	}
}
