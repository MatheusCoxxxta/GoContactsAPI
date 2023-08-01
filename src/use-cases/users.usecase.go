package usecases

import (
	"fmt"

	"api/src/db"
	"api/src/models"
)

func GetUsers(users []models.User) (u []models.User) {

	u = make([]models.User, 0)
	db.CONNECTION.Preload("Contact").Find(&u)

	return
}

func GetUser(user models.User, id string) (u models.User, err error) {
	u = user

	exec := db.CONNECTION.Preload("Contact").First(&u, id)

	if exec.Error != nil {
		err = exec.Error
	}

	return
}

func CreateUser(user *models.User) {
	db.CONNECTION.Create(&user)
}

func UpdateUser() {
	fmt.Println("Not implemented")
}

func DeleteUser(user *models.User, id string) (err error) {
	result := db.CONNECTION.First(&user, id)

	if result.Error != nil {
		err = result.Error
		return
	}

	db.CONNECTION.Delete(&user, id)

	return
}
