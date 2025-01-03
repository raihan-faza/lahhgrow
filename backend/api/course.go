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
	db_error := db.Create(&course).Error
	if db_error != nil {
		responses.BadRequest(c, "failed to create course")
		return
	}
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
	db_err := db.First(&course, course_id).Error
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

func DeleteCourse(c *gin.Context, db *gorm.DB) {
	var fail_message = "failed to delete course"
	course_id := c.Param("id")
	if course_id == "" {
		responses.BadRequest(c, fail_message)
		return
	}
	db_err := db.Delete(&models.Course{}, course_id).Error
	if db_err != nil {
		responses.BadRequest(c, fail_message)
		return
	}
	responses.GoodRequest(c, "course deleted succesfully")
	return
}

func GetCourses(c *gin.Context, db *gorm.DB) {
	var courses []models.Course
	db_err := db.Find(&courses).Error
	if db_err != nil {
		responses.BadRequest(c, "course unavailable")
	}
	c.JSON(200, courses)
	return
}

func GetCourse(c *gin.Context, db *gorm.DB) {
	var course models.Course
	course_id := c.Param("id")
	db_err := db.Find(&course, course_id).Error
	if db_err != nil {
		responses.BadRequest(c, "course unavailable")
	}
	c.JSON(200, course)
	return
}
