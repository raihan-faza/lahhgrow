package utils

import (
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
