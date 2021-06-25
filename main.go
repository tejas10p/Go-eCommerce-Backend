package main

import (
	"eCommerce/server"
	"log"
	"net/http"
)


func main() {
	router := server.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
}
