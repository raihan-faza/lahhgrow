package api

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/raihan-faza/lahhgrow/backend/models"
	"github.com/raihan-faza/lahhgrow/backend/utils"
	"gorm.io/gorm"
)

func CreateToken(c *gin.Context, db *gorm.DB, username string, password string, user_id string) (string, string, error) {
	var user models.Account
	hashed_password, hash_err := utils.CreateHash(password)
	if db.Find(&user, user_id).Error != nil {
		log.Fatal("user not found")
	}
	if hash_err != nil {
		panic(hash_err)
	}

	if hashed_password != user.Password {
		log.Fatal("somebody trying to do funny stuff")
		return "", "", fmt.Errorf("hahaha")
	}
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user_id,
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(), //ini limabelasmenit
	})
	signed_access_token, err := access_token.SignedString([]byte(utils.SecretKey))
	if err != nil {
		panic(err)
	}
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user_id,
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

func DecodeJwt(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return utils.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func GetJwtFromHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		log.Println("Authorization header not found")
		return ""
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		log.Println("Invalid Authorization header format")
		return ""
	}
	return parts[1]
}
