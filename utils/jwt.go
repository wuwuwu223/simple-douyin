package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/model"
)

var secretKey = []byte("secret key")

// generate a jwt token
func GenerateJwtToken(userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// parse jwt token
func ParseJwtToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//get userId from jwt token
func GetUserIdFromJwtToken(c *gin.Context) int64 {
	if userId, exists := c.Get("userId"); exists {
		return int64(userId.(float64))
	}
	return 0
}

//parse userId from jwt token
func GetUserIdFromToken(tokenString string) int64 {
	token, err := ParseJwtToken(tokenString)
	if err != nil {
		return 0
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int64(claims["userId"].(float64))
	}
	return 0
}

//jwt 中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Query("token")
		fmt.Println(tokenString)
		if tokenString == "" {
			c.JSON(http.StatusOK, gin.H{
				"Response": model.Response{StatusCode: 1, StatusMsg: "token is empty"},
			})
			c.Abort()
			return
		}
		token, err := ParseJwtToken(tokenString)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Response": model.Response{StatusCode: 1, StatusMsg: "token is invalid"},
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"])
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Response": model.Response{StatusCode: 1, StatusMsg: "token is invalid"},
			})
			c.Abort()
			return
		}
	}
}
