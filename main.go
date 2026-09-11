package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	env, err := loadEnv()
	if err != nil {
		fmt.Println(err)
		return
	}

	jar, _ := cookiejar.New(nil)
	client := http.Client{Jar: jar}

	resp, err := client.PostForm("http://localhost:8080/api/v2/auth/login", url.Values{
		"username": {env.User},
		"password": {env.Pass},
	})
	if err != nil {
		fmt.Print(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println(resp.Status)

}
