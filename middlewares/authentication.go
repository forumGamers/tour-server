package middlewares

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type User struct {
	Email			string
	Fullname		string
	Iat				int
	Id				int
	isVerified		bool
	PhoneNumber		string
	Username		string
}

func getSecretKey() (string,error){
	if err := godotenv.Load() ; err != nil {
		return "",errors.New("Internal Server Error")
	}

	return os.Getenv("SECRET"),nil
}

func Authentication(c *gin.Context){
	access_token := c.Request.Header.Get("access_token")

	if access_token == "" {
		panic("Invalid token")
	}

	secret,_ := getSecretKey()

	token,err := jwt.Parse(access_token,func(t *jwt.Token)(interface{},error){
		return []byte(secret),nil
	})

	if err != nil {
		panic("signature is invalid")
	}

	if !token.Valid {
		panic("Invalid token")
	}

	c.Set("user",token.Claims)

	c.Next()
}