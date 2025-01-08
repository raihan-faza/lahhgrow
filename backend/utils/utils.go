package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey string
var Snapc snap.Client

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

func CreateHash(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashed_password := string(bytes)
	return hashed_password, nil
}

func SetupMidtransConfig() {
	midtrans.ServerKey = os.Getenv("server_key")
	midtrans.ClientKey = os.Getenv("client_key")
}

func InitSnapDev() {
	Snapc.New(midtrans.ServerKey, midtrans.Sandbox)
}
