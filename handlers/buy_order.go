package handlers

import (
	"eCommerce/workers"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func BuyOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		ID := params["id"]
		result, err := workers.InitiatePurchase(ID)
		if err != nil {
			if err.Error() == "order retrieval fail" {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println("Cannot retrieve order")
			} else if err.Error() == "status check fail" {
				w.WriteHeader(http.StatusBadRequest)
				log.Println("Trying to purchase an already purchased or returned order")
			} else if err.Error() == "row scan fail" {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println("Error while scanning data rows")
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println("Could not update order status")
			}
			return
		}
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(result)
	}
}

// Handle 4xx Context
