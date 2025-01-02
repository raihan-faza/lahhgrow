package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raihan-faza/lahhgrow/backend/api"
	"gorm.io/gorm"
)

func MainRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.POST("/course", func(ctx *gin.Context) {
		api.CreateCourse(ctx, db)
	})
	return r
}
