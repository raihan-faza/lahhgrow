package models

import (
	"fmt"

	"github.com/go-co-op/gocron/v2"
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

type WalletCronManager struct {
	gorm.Model
	Scheduler *gocron.Scheduler
	WalletID  uint
	Wallet    Wallet
}

func (w *Wallet) AddBalance(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("you extorting me now?")
	}
	w.Balance = w.Balance + amount
	return nil
}

func (w *WalletCronManager) CutBalance(amount float64) error {
	if amount > w.Wallet.Balance {
		return fmt.Errorf("you ran out of money")
	}
	w.Wallet.Balance = w.Wallet.Balance - amount
	return nil
}
