package repositories

import (
	"api/src/db"
	"api/src/models"
)

func List() []models.User {

	u := make([]models.User, 0)
	db.CONNECTION.Preload("GithubAccount").Preload("Contact").Find(&u)

	return u
}

func FindById(id string) (u models.User, err error) {

	result := db.CONNECTION.Preload("Contact").Preload("GithubAccount").First(&u, id)

	if result.Error != nil {
		err = result.Error
	}

	return
}

func Create(user models.User) (u models.User) {

	db.CONNECTION.Create(&user)
	u = user

	return

}

func Update(mergedUser models.User) {

	db.CONNECTION.Save(&mergedUser)

}

func Delete(mergedUser models.User, id string) {

	db.CONNECTION.Delete(&mergedUser, id)

}
