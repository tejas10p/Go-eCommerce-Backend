package handlers

import (
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for _, v := range database.Products {
			if v.ID == params["id"] {
				json.NewEncoder(w).Encode(v)
				return
			}
		}
		json.NewEncoder(w).Encode(models.Product{})
	}
}