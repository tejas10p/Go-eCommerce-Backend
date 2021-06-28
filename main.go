package main

import (
	"eCommerce/database"
	"eCommerce/server"
	"log"
	"net/http"
)

func main() {
	database.Init()
	defer database.Close()
	router := server.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
}
