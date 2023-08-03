package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"api/src/models"
)

func GetGithubByUsername(username string) (githubAccount models.GithubAccount) {
	url := "https://api.github.com/users/" + username
	response, _ := http.Get(url)

	responseBytes, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(responseBytes, &githubAccount)

	return
}
