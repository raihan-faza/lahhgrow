package api

import (
	"github.com/gin-gonic/gin"
	"github.com/raihan-faza/lahhgrow/backend/models"
	"github.com/raihan-faza/lahhgrow/backend/responses"
	"gorm.io/gorm"
)

// blom di save ke db, blom bikin initializer dbnya
func CreateCourse(c *gin.Context, db *gorm.DB) {
	var course models.Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		responses.BadRequest(c, err)
		return
	}
	db.Create(&course)
	responses.GoodRequest(c, "course created")
	return
}

func UpdateCourse(c *gin.Context) {
	var course models.Course
	err := c.ShouldBindJSON(&course)
	if err != nil {
		responses.BadRequest(c, err)
		return
	}
	responses.GoodRequest(c, "course updated")
	return
}
