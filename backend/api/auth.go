package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/raihan-faza/lahhgrow/backend/utils"
)

func CreateToken(c *gin.Context, username string, user_id string) (string, string, error) {
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user_id,
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(), //ini limabelasmenit
	})
	signed_access_token, err := access_token.SignedString([]byte(utils.SecretKey))
	if err != nil {
		panic(err)
	}
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user_id,
		"username": username,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(), //ini 7hari
	})
	signed_refresh_token, err := refresh_token.SignedString([]byte(utils.SecretKey))
	if err != nil {
		panic(err)
	}
	return signed_access_token, signed_refresh_token, nil
}

func RefreshToken(c *gin.Context, refresh_token string) {
	return
}
