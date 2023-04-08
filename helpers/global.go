package helpers

import (
	"errors"
	"regexp"
	"strings"

	md "github.com/forumGamers/tour-service/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateInput(data map[string]string) error{
	for key, value := range data {
		if !strings.Contains(strings.ToLower(key),"password") && ValidateInvalidCharacter(value){
			return errors.New("Invalid data")
		}
	}

	return nil
}

func ValidateInvalidCharacter(data string) bool {
	return regexp.MustCompile(`[^a-zA-Z0-9.,\-\s@]`).MatchString(data)
}

func GetUser(c *gin.Context) md.User {
	claim := c.MustGet("user").(jwt.MapClaims)
	var user md.User

	for key, val := range claim {
		switch key {
		case "email":
			user.Email = val.(string)
		case "fullName":
			user.Fullname = val.(string)
		case "iat":
			user.Iat = int(val.(float64))
		case "id":
			user.Id = int(val.(float64))
		case "isVerified":
			user.IsVerified = val.(bool)
		case "phoneNumber":
			user.PhoneNumber = val.(string)
		case "username":
			user.Username = val.(string)
		}
	}
	return user
}