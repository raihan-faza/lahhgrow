package api

import (
	"github.com/gin-gonic/gin"
	"github.com/raihan-faza/lahhgrow/backend/models"
	"github.com/raihan-faza/lahhgrow/backend/responses"
	"gorm.io/gorm"
)

func CreateCourse(c *gin.Context, db *gorm.DB) {
	var course models.Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		responses.BadRequest(c, "failed to create course")
		return
	}
	db.Create(&course)
	responses.GoodRequest(c, "course created")
	return
}

func UpdateCourse(c *gin.Context, db *gorm.DB) {
	var course models.Course
	var input models.Course
	var fail_message = "failed to update course"
	course_id := c.DefaultQuery("id", "")
	if course_id == "" {
		responses.BadRequest(c, fail_message)
		return
	}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		responses.BadRequest(c, fail_message)
		return
	}
	db_err := db.First(&course, course_id)
	if db_err != nil {
		responses.BadRequest(c, fail_message)
		return
	}
	update_data := db.Model(&course).Updates(&input)
	if update_data.Error != nil {
		responses.BadRequest(c, fail_message)
	}
	responses.GoodRequest(c, "course updated succesfully")
	return
}
