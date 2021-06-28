package models

type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber int64  `json:"phone"`
	Address     string `json:"address"`
}
