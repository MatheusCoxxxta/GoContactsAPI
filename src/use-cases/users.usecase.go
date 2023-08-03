package usecases

import (
	mp "github.com/geraldo-labs/merge-struct"

	"api/src/controllers/providers"
	"api/src/models"
	"api/src/repositories"
)

func GetUsers() []models.User {

	return repositories.List()

}

func GetUser(id string) (u models.User, err error) {

	u, err = repositories.FindById(id)

	return
}

func CreateUser(user models.User) (u models.User) {

	githubAccount := providers.GetGithubByUsername(user.GithubUsername)

	user.GithubAccount = githubAccount

	u = repositories.Create(user)

	return
}

func UpdateUser(id string, modified models.User) (u models.User, err error) {

	u, repoErr := repositories.FindById(id)

	if repoErr != nil {
		err = repoErr
		return
	}

	_, err = mp.Struct(&u, modified)

	repositories.Update(u)

	return
}

func DeleteUser(id string) (err error) {

	u, err := repositories.FindById(id)

	repositories.Delete(u, id)

	return
}
