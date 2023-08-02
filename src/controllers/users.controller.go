package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var bodyUser models.User
	json.NewDecoder(r.Body).Decode(&bodyUser)

	modified := usecases.UpdateUser(params["id"], bodyUser)

	json.NewEncoder(w).Encode(modified)
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
