package usecases

import (
	"fmt"

	"api/src/db"
	"api/src/models"
)

func GetUsers(users *[]models.User) {
	db.CONNECTION.Preload("Contact").Find(&users)
}

func GetUser(user *models.User, id string) (err error) {
	result := db.CONNECTION.Preload("Contact").First(&user, id)

	if result.Error != nil {
		err = result.Error
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
