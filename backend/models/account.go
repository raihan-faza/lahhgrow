package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Address     string `json:"address" binding:"required"`
	City        string `json:"city" binding:"required"`
	Postcode    string `json:"postcode" binding:"required"`
	CountryCode string `json:"country_code" binding:"required"`
	Wallet      Wallet
}
