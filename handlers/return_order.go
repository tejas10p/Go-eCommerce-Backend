package handlers

import (
	"eCommerce/database"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func ReturnOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatalf("Invalid ID - %s", err.Error())
		}
		json.NewEncoder(w).Encode(database.ReturnOrder(ID))
	}
}
