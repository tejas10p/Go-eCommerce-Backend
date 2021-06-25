package handlers

import (
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"net/http"
)

func ValidateUser(order models.Order) bool {
	for _, v := range database.Users {
		if v.Email == order.Customer.Email {
			return true
		}
	}
	return false
}

func BuyProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var purchase models.Order
		database.Orders = append(database.Orders, purchase)
		_ = json.NewDecoder(r.Body).Decode(&purchase)
		var modifiedProducts []models.Product
		if !ValidateUser(purchase) {
			json.NewEncoder(w).Encode("Customer not found")
			return
		}
		for i, v := range database.Products {
			for j, pur := range purchase.ProductId {
				if v.ID == pur {
					if v.Quantity >= purchase.Quantity[j] {
						database.Products[i].Quantity -= purchase.Quantity[j]
						modifiedProducts = append(modifiedProducts, database.Products[i])
						break
					}
				}
			}
		}
		json.NewEncoder(w).Encode(modifiedProducts)
	}
}
