package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}
