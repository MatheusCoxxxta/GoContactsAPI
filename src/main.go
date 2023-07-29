package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"api/src/controllers"
	"api/src/db"
)

func initializeRouter() {
	r := mux.NewRouter()

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
