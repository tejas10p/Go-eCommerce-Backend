package handlers

import (
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"net/http"
)

func GetUserOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user)
		var requiredOrders []models.Order
		for _, v := range database.Orders {
			if v.Customer.Email == user.Email {
				requiredOrders = append(requiredOrders, v)
			}
		}
		json.NewEncoder(w).Encode(requiredOrders)
	}
}
