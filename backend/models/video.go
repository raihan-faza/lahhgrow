package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Name string `json:"name"`
	Path string `json:"path"`
	Size int    `json:"size"`
}
