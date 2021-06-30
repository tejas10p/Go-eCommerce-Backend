package database

import (
	"fmt"
	"log"
)

func RunDatabaseMigrations() {
	_, err = Db.Exec(
		"CREATE TABLE IF NOT EXISTS user(name varchar(50) NOT NULL, email varchar(50) NOT NULL, phone bigint, Address varchar(250), PRIMARY KEY(email))")
	if err != nil {
		log.Fatalf("Error creating table user - %s", err.Error())
		return
	}
	_, err = Db.Exec(
		"CREATE TABLE IF NOT EXISTS product(id int AUTO_INCREMENT, name varchar(50) NOT NULL, price int NOT NULL, category varchar(50), PRIMARY KEY(id))")
	if err != nil {
		log.Fatalf("Error creating table product - %s", err.Error())
		return
	}
	_, err = Db.Exec(
		"CREATE TABLE IF NOT EXISTS orders(id int AUTO_INCREMENT, email varchar(50) NOT NULL, address varchar(250), status varchar(20) NOT NULL, PRIMARY KEY(id), FOREIGN KEY (email) REFERENCES user(email))")
	if err != nil {
		log.Fatalf("Error creating table order - %s", err.Error())
		return
	}
	_, err = Db.Exec(
		"CREATE TABLE IF NOT EXISTS item(orderid int NOT NULL, productid int NOT NULL, quantity int NOT NULL, amount int NOT NULL, FOREIGN KEY (orderid) REFERENCES orders(id), FOREIGN KEY (productid) REFERENCES product(id))")
	if err != nil {
		log.Fatalf("Error creating table item - %s", err.Error())
		return
	}
	fmt.Println("Successfully created tables!")
}

func RollBackDatabaseMigrations() {
	_, err = Db.Exec("DROP TABLE product")
	if err != nil {
		log.Fatalf("Error while dropping product table - %s", err.Error())
		return
	}
	_, err = Db.Exec("DROP TABLE user")
	if err != nil {
		log.Fatalf("Error while dropping product user - %s", err.Error())
		return
	}
	_, err = Db.Exec("DROP TABLE orders")
	if err != nil {
		log.Fatalf("Error while dropping product orders - %s", err.Error())
		return
	}
	_, err = Db.Exec("DROP TABLE item")
	if err != nil {
		log.Fatalf("Error while dropping product item - %s", err.Error())
		return
	}
}
