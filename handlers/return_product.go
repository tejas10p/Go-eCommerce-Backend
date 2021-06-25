package handlers

import (
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"net/http"
)

func ReturnProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var returns models.Order
		database.Orders = append(database.Orders, returns)
		_ = json.NewDecoder(r.Body).Decode(&returns)
		var modifiedProducts []models.Product
		if !ValidateUser(returns) {
			json.NewEncoder(w).Encode("Customer not found")
			return
		}
		for i, v := range database.Products {
			for j, pur := range returns.ProductId {
				if v.ID == pur {
					database.Products[i].Quantity += returns.Quantity[j]
					modifiedProducts = append(modifiedProducts, database.Products[i])
					break
				}
			}
		}
		json.NewEncoder(w).Encode(modifiedProducts)
	}
}
