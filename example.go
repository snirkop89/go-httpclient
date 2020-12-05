package main

import (
	"fmt"

	"github.com/snirkop89/go-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {

	builder := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5)

	client := builder.Build()

	// commonHeaders := make(http.Header)
	// commonHeaders.Set("Authorization", "Bearer ABC-123")

	return client
}

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			getUrls()
		}()
	}
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getUrls() {
	// headers := make(http.Header)
	// headers.Set("Authorization", "Bearer ABC-123")

	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	var user User
	if err := response.UnmarshalJson(&user); err != nil {
		panic(err)
	}
	// fmt.Println(user.FirstName)
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode())

	fmt.Println(response.String())
}
