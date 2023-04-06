package cmd

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	cfg "github.com/forumGamers/tour-service/config"
	h "github.com/forumGamers/tour-service/helpers"
	md "github.com/forumGamers/tour-service/middlewares"
	m "github.com/forumGamers/tour-service/models"
	"github.com/gin-gonic/gin"
)

func CreateTeam(c *gin.Context){
	s,_ := c.Get("user")

	user,_ := s.(md.User)

	name := c.PostForm("name")

	players := c.PostFormArray("player")

	if name == "" {
		panic("Invalid data")
	}

	if h.ValidateInvalidCharacter(name) {
		panic("name do not allow contains symbol")
	}

	var list []int

	for _, player := range players {
		playerId,err := strconv.Atoi(player)

		if err != nil {
			panic("Invalid data")
		}

		list = append(list, playerId)
	}

	image,err := c.FormFile("image")

	if err != nil {
		panic(err.Error())
	}

	if err := c.SaveUploadedFile(image,"uploads/"+image.Filename) ; err != nil {
		panic(err.Error())
	}

	file,_ := os.Open("uploads/"+image.Filename)

	data,err := ioutil.ReadAll(file)

	if err != nil {
		panic(err.Error())
	}

	urlCh := make(chan string)
	fileIdCh := make(chan string)
	errUpload := make(chan error)
	errCh := make(chan error)

	go func(data []byte,imageName string){
		url,fileId,err := cfg.UploadImage(data,imageName,"teamImage")

		if err != nil {
			errUpload <- errors.New("Bad Gateway")
			urlCh <- ""
			fileIdCh <- ""
			return
		}

		errUpload <- nil
		urlCh <- url
		fileIdCh <- fileId
	}(data,image.Filename)

	go func(name string,user md.User,list []int) {
		if err := <- errUpload ; err != nil {
			errCh <- err
			return
		}

		if _,err := getDb().Collection("team").InsertOne(context.Background(),m.Team{
			Name: name,
			OwnerId: user.Id,
			UserId: list,
			Image: <- urlCh,
			ImageId: <- fileIdCh,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}) ; err != nil {
			errCh <- err
			return
		}

		errCh <- nil
	}(name,user,list)

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusCreated,gin.H{"message":"success"})
}