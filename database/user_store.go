package database

import (
	"database/sql"
	"eCommerce/models"
	"log"
)

func GetUsers() []models.User {
	result, err := Db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatalf("Cannot get users - %s", err.Error())
	}
	var users []models.User
	for result.Next() {
		var currentUser models.User
		err = result.Scan(&currentUser.Name, &currentUser.Email, &currentUser.PhoneNumber, &currentUser.Address)
		if err != nil {
			log.Fatalf("Error while scanning data rows - %s", err.Error())
		}
		users = append(users, currentUser)
	}
	return users
}

func AddUser(user models.User) models.User {
	_, err := Db.Exec("INSERT INTO user VALUES (?, ?, ?, ?)", user.Name, user.Email, user.PhoneNumber, user.Address)
	if err != nil {
		log.Fatalf("Cannot add product - %s", err.Error())
	}
	return user
}

func GetUserOrders(email string) []models.Item {
	var items []models.Item
	result, err := Db.Query("SELECT id FROM orders WHERE email = ?", email)
	for result.Next() {
		var currentId string
		err = result.Scan(&currentId)
		if err != nil {
			log.Fatalf("Error getting order ids - %s", err.Error())
		}
		var itemResult *sql.Rows
		itemResult, err = Db.Query("SELECT * FROM item WHERE orderid = ?", currentId)
		if err != nil {
			log.Fatalf("Error retrieving order items - %s", err.Error())
		}
		for itemResult.Next() {
			var currentItem models.Item
			err = result.Scan(&currentItem.OrderId, &currentItem.ProductId, &currentItem.Quantity, &currentItem.Amount)
			if err != nil {
				log.Fatalf("Error while scanning data rows - %s", err.Error())
			}
			items = append(items, currentItem)
		}
	}
	return items
}
