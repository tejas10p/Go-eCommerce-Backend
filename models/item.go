package models

type Item struct {
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
	Amount    int `json:"amount"`
}
