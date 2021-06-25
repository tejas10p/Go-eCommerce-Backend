package models

type Order struct {
	ProductId []string `json:"productid"`
	Quantity []int `json:"quantity"`
	Customer *User `json:"customer"`
}

// 1 -> 2 2 -> 4 === [1, 2] === [2, 4] (Club Id and Quantity)