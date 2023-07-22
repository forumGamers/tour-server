package helpers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"regexp"
	"strings"

	cfg "github.com/forumGamers/tour-service/config"
	m "github.com/forumGamers/tour-service/models"
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

func GetUser(c *gin.Context) m.User {
	claimMap,ok := c.Get("user")

	var user m.User

	if !ok {
		return user
	}

	claim,oke := claimMap.(jwt.MapClaims)

	if !oke {
		return user
	}

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
		case "StoreId" :
			user.StoreId = int(val.(float64))
		case "role" :
			user.Role = val.(string)
		case "point" :
			user.Point = int(val.(float64))
		case "experience" :
			user.Exp = int(val.(float64))
		}
	}
	return user
}

func UploadImage(data []byte,imageName string,folderName string,urlCh chan string,fileIdCh chan string,errCh chan error){
	url,fileId,err := cfg.UploadImage(data,imageName,folderName)

		if err != nil {
			errCh <- errors.New("Bad Gateway")
			urlCh <- ""
			fileIdCh <- ""
			return
		}

		errCh <- nil
		urlCh <- url
		fileIdCh <- fileId
}

func IsImage(file *multipart.FileHeader) error {
	imgExt := []string{"png", "jpg", "jpeg", "gif", "bmp"}

	ext := strings.ToLower(strings.TrimPrefix(strings.TrimPrefix(file.Filename,"."),"."))

	for _,val := range imgExt {
		if val == ext {
			if file.Size > 10*1024*1024 {
				return fmt.Errorf("file cannot be larger than 10 mb")
			}

			return nil
		}
	}

	return fmt.Errorf("file type is not supported")
}