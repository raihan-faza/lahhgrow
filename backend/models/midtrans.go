package models

import "gorm.io/gorm"

type TransactionDetails struct {
	gorm.Model
	OrderId           string
	GrossAmount       string
	CustomerDetailsID int
}

type CustomerDetails struct {
	gorm.Model
	FName              string
	LName              string
	Phone              string
	Address            string
	City               string
	Postcode           string
	CountryCode        string
	TransactionDetails []TransactionDetails
}

// i donno what am i doing smh
