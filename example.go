package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/snirkop89/go-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	// commonHeaders := make(http.Header)
	// commonHeaders.Set("Authorization", "Bearer ABC-123")

	// client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()
	getUrls()
}

type User struct {
	FirstName string `json:first_name`
	LastName  string `json:last_name`
}

func getUrls() {
	headers := make(http.Header)
	// headers.Set("Authorization", "Bearer ABC-123")

	response, err := githubHttpClient.Get("https://api.github.com", headers)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
