package database

import (
	"database/sql"
	"eCommerce/kafka"
	"eCommerce/models"
	"errors"
	"log"
	"strconv"
)

func AddOrder(order models.Order) models.Order {
	order.Status = "Created"
	_, err := Db.Exec("INSERT INTO orders (email, address, status) VALUES (?, ?, ?, ?)", order.UserEmail, order.Address, order.Status)
	if err != nil {
		log.Fatalf("Cannot add order - %s", err.Error())
	}
	var result *sql.Rows
	result, err = Db.Query("SELECT last_insert_id()")
	if err != nil {
		log.Fatalf("Error retrieving ID for product - %s", err.Error())
	}
	result.Scan(&order.OrderId)
	return order
}

func BuyOrder() (int, error) {
	c := kafka.NewConsumer("orders")
	kafkaResult := kafka.ConsumeMessage(c, "purchases")
	kafka.CloseConsumer(c)
	id, err := strconv.Atoi(kafkaResult)
	if err != nil {
		return 0, errors.New("order retrieval fail")
	}
	result, err := Db.Query("SELECT * FROM item WHERE orderid = ?", id)
	if err != nil {
		return 0, errors.New("order retrieval fail")
	}
	if checkStatus(id, "Created") {
		return 0, errors.New("status check fail")
	}
	totalAmount := 0
	for result.Next() {
		var currentItem models.Item
		err = result.Scan(&currentItem.OrderId, &currentItem.ProductId, &currentItem.Quantity, &currentItem.Amount)
		if err != nil {
			return 0, errors.New("row scan fail")
		}
		totalAmount += currentItem.Amount
	}
	_, err = Db.Exec("UPDATE orders SET status = ? WHERE id = ?", "Purchased", id)
	if err != nil {
		return 0, errors.New("status update fail")
	}
	return totalAmount, nil
}

func ReturnOrder(id int) int {
	result, err := Db.Query("SELECT * FROM item WHERE orderid = ?", id)
	if err != nil {
		log.Fatalf("Cannot retrieve order - %s", err.Error())
	}
	if checkStatus(id, "Purchased") {
		log.Fatalf("Trying to return an unpurchased or already returned order")
	}
	totalAmount := 0
	for result.Next() {
		var currentItem models.Item
		err = result.Scan(&currentItem.OrderId, &currentItem.ProductId, &currentItem.Quantity, &currentItem.Amount)
		if err != nil {
			log.Fatalf("Error while scanning data rows - %s", err.Error())
		}
		totalAmount += currentItem.Amount
	}
	_, err = Db.Exec("UPDATE orders SET status = ? WHERE id = ?", "Returned", id)
	if err != nil {
		log.Fatalf("Could not update order status - %s", err.Error())
	}
	return totalAmount
}
