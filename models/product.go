package models

type Product struct {
	ProductId int    `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Category  string `json:"category"`
}
