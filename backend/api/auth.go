package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(c *gin.Context, username string, user_id string, secret_key int) (string, string, error) {
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user_id,
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(), //ini limabelasmenit
	})
	signed_access_token, err := access_token.SignedString(secret_key)
	if err != nil {
		panic(err)
	}
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user_id,
		"username": username,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(), //ini 7hari
	})
	signed_refresh_token, err := refresh_token.SignedString(secret_key)
	if err != nil {
		panic(err)
	}
	return signed_access_token, signed_refresh_token, nil
}
