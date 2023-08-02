package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"api/src/db"
	"api/src/models"
	usecases "api/src/use-cases"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users := usecases.GetUsers()

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user, err := usecases.GetUser(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	u := usecases.CreateUser(user)

	json.NewEncoder(w).Encode(u)
}

// TODO: refactor UpdateUser to use an usecase.UpdateUser

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user models.User

	result := db.CONNECTION.First(&user, params["id"])

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return
	}

	json.NewDecoder(r.Body).Decode(&user)

	db.CONNECTION.Save(&user)

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := usecases.DeleteUser(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode("The user was deleted successfully")
}
