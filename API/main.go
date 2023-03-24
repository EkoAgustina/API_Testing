package main

import (
	"net/http"
	"fmt"

	"github.com/jeypc/go-jwt-mux/middlewares"

	"github.com/jeypc/go-jwt-mux/controllers/authcontroller"
	"github.com/jeypc/go-jwt-mux/controllers/productcontroller"
	"github.com/jeypc/go-jwt-mux/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/addProduct", productcontroller.AddProduct).Methods("POST")
	api.HandleFunc("/products", productcontroller.GetProduct).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	
	fmt.Print(http.ListenAndServe(":8081", r))






}
