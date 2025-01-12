package models

import (
	"fmt"
	"time"

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
	Scheduler gocron.Scheduler
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

func (w *WalletCronManager) StartCutBalance() error {
	var amount = float64(1000)
	if amount > w.Wallet.Balance {
		return fmt.Errorf("you ran out of money")
	}
	w.Scheduler.NewJob(
		gocron.DurationJob(1*time.Minute),
		gocron.NewTask(func() {
			w.Wallet.Balance = w.Wallet.Balance - amount
		}),
	)
	w.Scheduler.Start()
	return nil
}

func (w *WalletCronManager) StopCutBalance() error {
	w.Scheduler.StopJobs()
	return nil
}
