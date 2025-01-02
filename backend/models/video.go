package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Path     string `json:"path"`
	Size     uint   `json:"size"`
	CourseID uint   `json:"course_id" binding:"required"`
}
