package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Balance   float64
	AccountID uint `json:"account_id" binding:"required"`
}

type AddBalanceRequest struct {
	Balance float64 `json:"balance" binding:"required"`
}

func (w *Wallet) AddBalance(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("you extorting me now?")
	}
	w.Balance = w.Balance + amount
	return nil
}
