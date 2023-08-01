package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"api/src/controllers"
	"api/src/db"
	"api/src/middlewares"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.Use(middlewares.JsonMiddleware)

	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	db.InitializeDatabase()
	initializeRouter()
}
