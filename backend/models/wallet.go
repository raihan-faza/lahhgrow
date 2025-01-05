package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Balance   float64
	AccountID uint `json:"account_id" binding:"required"`
}
