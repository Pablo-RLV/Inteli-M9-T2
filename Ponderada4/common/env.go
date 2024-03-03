package common

import (
	"fmt"
	godotenv "github.com/joho/godotenv"
)

func LoadEnv(){
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
}
