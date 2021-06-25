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
		if !newUser.ValidateEmail() {
			json.NewEncoder(w).Encode(models.User{})
			return
		}
		database.Users = append(database.Users, newUser)
		json.NewEncoder(w).Encode(newUser)
	}
}
