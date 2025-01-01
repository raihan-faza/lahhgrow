package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name   string  `json:"name"`
	Info   string  `json:"info"`
	Videos []Video `json:"videos"`
}
