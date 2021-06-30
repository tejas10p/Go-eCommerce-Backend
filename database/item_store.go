package database

import (
	"eCommerce/models"
	"log"
)

func checkStatus(id int, requiredStatus string) bool {
	result, err := Db.Query("SELECT status FROM orders WHERE id = ?", id)
	if err != nil {
		log.Fatalf("Could not get order status - %s", err.Error())
	}
	var currentStatus string
	result.Next()
	result.Scan(&currentStatus)
	return currentStatus != requiredStatus
}

func AddItem(item models.Item) models.Item {
	if checkStatus(item.OrderId, "Created") {
		log.Fatalf("Attempting to add products in already purchased or returned order")
	}
	_, err = Db.Exec("INSERT INTO user VALUES (?, ?, ?, ?)", item.OrderId, item.ProductId, item.Quantity, item.Amount)
	if err != nil {
		log.Fatalf("Cannot add item - %s", err.Error())
	}
	return item
}
