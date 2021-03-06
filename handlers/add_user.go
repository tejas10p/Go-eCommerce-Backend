package handlers

import (
	"eCommerce/database"
	"eCommerce/models"
	"encoding/json"
	"net/http"
)

func AddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var newUser models.User
		_ = json.NewDecoder(r.Body).Decode(&newUser)
		json.NewEncoder(w).Encode(database.AddUser(newUser))
	}
}
