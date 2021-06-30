package handlers

import (
	"eCommerce/database"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUserOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		json.NewEncoder(w).Encode(database.GetUserOrders(params["email"]))
	}
}
