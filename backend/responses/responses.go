package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, err string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": err,
	})
}

func GoodRequest(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
