package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB
var err error

func Init() {
	Db, err = sql.Open("mysql", "root:MySQL@10@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatalf("Cannot open database due to some error - %s", err.Error())
		return
	}
	fmt.Println("Successfully connected to database server!")

	_, err = Db.Exec("CREATE DATABASE IF NOT EXISTS eCommerce")
	if err != nil {
		log.Fatalf("Error creating database - %s", err.Error())
		return
	}
	_, err = Db.Exec("USE eCommerce")
	if err != nil {
		log.Fatalf("Error selcting database - %s", err.Error())
		return
	}
	fmt.Println("Successfully connected to database!")

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

func Close() {
	err = Db.Close()
	if err != nil {
		log.Fatalf("Cannot close database due to error - %s", err.Error())
	}
	fmt.Println("Connection to database closed!")
}
