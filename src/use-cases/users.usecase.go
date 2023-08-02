package usecases

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	mp "github.com/geraldo-labs/merge-struct"

	"api/src/db"
	"api/src/models"
)

func GetUsers() (u []models.User) {

	u = make([]models.User, 0)
	db.CONNECTION.Preload("Contact").Find(&u)

	return
}

func GetUser(id string) (u models.User, err error) {
	u = models.User{}

	result := db.CONNECTION.Preload("Contact").Preload("GithubAccount").First(&u, id)

	if result.Error != nil {
		err = result.Error
	}

	return
}

func CreateUser(user models.User) (u models.User) {

	url := "https://api.github.com/users/" + user.GithubUsername
	response, _ := http.Get(url)

	responseBytes, _ := ioutil.ReadAll(response.Body)

	var githubAccount models.GithubAccount
	json.Unmarshal(responseBytes, &githubAccount)

	user.GithubAccount = githubAccount

	u = user
	db.CONNECTION.Create(&u)

	return
}

func UpdateUser(id string, modified models.User) (u models.User) {

	result := db.CONNECTION.First(&u, id)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return
	}

	_, err := mp.Struct(&u, modified)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db.CONNECTION.Save(&u)

	return
}

func DeleteUser(id string) (err error) {
	u := models.User{}

	result := db.CONNECTION.First(&u, id)

	if result.Error != nil {
		err = result.Error
		return
	}

	db.CONNECTION.Delete(&u, id)

	return
}
