package server

import (
	"eCommerce/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/products", handlers.GetProducts()).Methods("GET")
	router.HandleFunc("/products/{id}", handlers.GetProduct()).Methods("GET")
	router.HandleFunc("/products", handlers.AddProduct()).Methods("POST")
	router.HandleFunc("/products/{id}", handlers.DeleteProduct()).Methods("DELETE")
	router.HandleFunc("/products/{id}", handlers.UpdateProduct()).Methods("PUT")
	router.HandleFunc("/orders/buy/{id}", handlers.BuyOrder()).Methods("PUT")
	router.HandleFunc("/orders/return/{id}", handlers.ReturnOrder()).Methods("PUT")
	router.HandleFunc("/users", handlers.GetUsers()).Methods("GET")
	router.HandleFunc("/users", handlers.AddUser()).Methods("POST")
	router.HandleFunc("/users/orderList/{email}", handlers.GetUserOrders()).Methods("GET")
	router.HandleFunc("/orders", handlers.CreateOrder()).Methods("POST")
	router.HandleFunc("/orders/item", handlers.AddItem()).Methods("POST")
	return router
}

//Buy Return on Order, API to see order of users
