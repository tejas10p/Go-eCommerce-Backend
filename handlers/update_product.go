package handlers

import (
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatalf("Invalid ID - %s", err.Error())
		}
		var newProduct models.Product
		_ = json.NewDecoder(r.Body).Decode(&newProduct)
		json.NewEncoder(w).Encode(database.UpdateProduct(ID, newProduct))
	}
}
