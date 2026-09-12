package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Pass string
	User string
}

func loadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	eC := &Config{
		Pass: os.Getenv("qbPass"),
		User: os.Getenv("qbUsr"),
	}

	return eC, nil
}

func qBitAuth(client *http.Client) {
	env, err := loadEnv()

	resp, err := client.PostForm("http://localhost:8080/api/v2/auth/login", url.Values{
		"username": {env.User},
		"password": {env.Pass},
	})
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(resp.Status)
	defer resp.Body.Close()
}
