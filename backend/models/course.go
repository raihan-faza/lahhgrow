package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name   string   `json:"name" binding:"required"`
	Info   string   `json:"info" binding:"required"`
	Videos *[]Video `json:"videos"`
}
