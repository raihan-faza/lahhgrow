package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/raihan-faza/lahhgrow/backend/responses"
)

func file_exist(file_name string) error {
	_, err := os.Stat(file_name)
	if err != nil {
		return err
	}
	return nil
}

func GetVideo(c *gin.Context, file_name string) {
	err := file_exist(file_name)
	if err != nil {
		responses.BadRequest(c, err)
		return
	}
	responses.GoodRequest(c, file_name)
	return
}
