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
	router.HandleFunc("/products/buy", handlers.BuyProduct()).Methods("PUT")
	router.HandleFunc("/products/return", handlers.ReturnProduct()).Methods("PUT")
	router.HandleFunc("/users", handlers.GetUsers()).Methods("GET")
	router.HandleFunc("/users", handlers.AddUser()).Methods("POST")
	router.HandleFunc("/users/orderList", handlers.GetUserOrders()).Methods("GET")
	return router
}
//Buy Return on Order, API to see order of users