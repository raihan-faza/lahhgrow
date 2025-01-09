package models

import "gorm.io/gorm"

type TransactionDetails struct {
	gorm.Model
	OrderId           string
	GrossAmount       string
	CustomerDetailsID int
	Status            string
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

type TopupRequest struct {
	GrossAmount int
}
