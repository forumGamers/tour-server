package cmd

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	h "github.com/forumGamers/tour-service/helpers"
	m "github.com/forumGamers/tour-service/models"
	v "github.com/forumGamers/tour-service/validation"
	"github.com/gin-gonic/gin"
)

func CreateTeam(c *gin.Context){

	user := h.GetUser(c)

	name := c.PostForm("name")

	description := c.PostForm("description")

	players := c.PostFormArray("player")

	if err := v.ValidateCreateTeam(name,description,players) ; err != nil {
		panic(err.Error())
	}

	var list []int

	for _, player := range players {
		playerId,err := strconv.Atoi(player)

		if err != nil {
			panic("Invalid data")
		}

		list = append(list, playerId)
	}

	list = append(list, user.Id)

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

	go h.UploadImage(data,image.Filename,"teamImage",urlCh,fileIdCh,errUpload)

	go func(name string,user m.User,list []int,description string) {
		if err := <- errUpload ; err != nil {
			errCh <- err
			return
		}

		if _,err := getDb().Collection("team").InsertOne(context.Background(),m.Team{
			Name: name,
			OwnerId: user.Id,
			UserId: list,
			Description: description,
			Image: <- urlCh,
			ImageId: <- fileIdCh,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}) ; err != nil {
			errCh <- err
			return
		}

		errCh <- nil
	}(name,user,list,description)

	file.Close()

	os.Remove("uploads/"+image.Filename)

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusCreated,gin.H{"message":"success"})
}