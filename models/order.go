package models

type Order struct {
	OrderId   int    `json:"order_id"`
	UserEmail string `json:"email"`
	Address   string `json:"address"`
	Status    string `json:"status"`
}
