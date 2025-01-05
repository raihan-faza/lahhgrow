package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Wallet   Wallet
}
