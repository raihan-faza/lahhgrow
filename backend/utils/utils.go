package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SecretKey string

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	SecretKey := os.Getenv("secret_key")
	if SecretKey == "" {
		log.Fatal("secret key not found")
	}
}
