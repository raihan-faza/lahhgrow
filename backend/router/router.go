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
	r.PUT("/course/:id", func(ctx *gin.Context) {
		api.UpdateCourse(ctx, db)
	})
	r.DELETE("/course/:id", func(ctx *gin.Context) {
		api.DeleteCourse(ctx, db)
	})
	r.GET("/course", func(ctx *gin.Context) {
		api.GetCourses(ctx, db)
	})
	r.GET("/course/:id", func(ctx *gin.Context) {
		api.GetCourse(ctx, db)
	})
	r.POST("/account", func(ctx *gin.Context) {
		api.CreateAccount(ctx, db)
	})
	r.DELETE("/account/:id", func(ctx *gin.Context) {
		api.DeleteAccount(ctx, db)
	})
	r.GET("/wallet/:id", func(ctx *gin.Context) {
		api.GetBalance(ctx, db)
	})
	r.POST("/wallet/:id", func(ctx *gin.Context) {
		api.AddBalance(ctx, db)
	})
	return r
}
