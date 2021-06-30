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

	RunDatabaseMigrations()
}

func Close() {
	err = Db.Close()
	if err != nil {
		log.Fatalf("Cannot close database due to error - %s", err.Error())
	}
	fmt.Println("Connection to database closed!")
}
