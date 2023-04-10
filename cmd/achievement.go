package cmd

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	h "github.com/forumGamers/tour-service/helpers"
	m "github.com/forumGamers/tour-service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAchievement(c *gin.Context){

	gameId := c.Param("gameId")

	name := c.PostForm("name")

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

	id,err := primitive.ObjectIDFromHex(gameId)

	if err != nil {
		panic("Invalid ObjectID")
	}

	urlCh := make(chan string)
	fileIdCh := make(chan string)
	errUpload := make(chan error)
	errCh := make(chan error)

	go h.UploadImage(data,image.Filename,"achievementImage",urlCh,fileIdCh,errUpload)

	go func(name string,gameId string,id primitive.ObjectID) {

		if err := <- errUpload ; err != nil {
			errCh <- err
			return
		}

		if _,err := getDb().Collection("achievement").InsertOne(context.Background(),m.Achievement{
			Name: name,
			Image: <- urlCh,
			ImageId: <- fileIdCh,
			GameId: id,
			CreatedAt: time.Now(),
			UpdatedAt:time.Now(),
		}) ; err != nil {
			errCh <- err
			return
		}

		errCh <- nil
	}(name,gameId,id)

	file.Close()

	os.Remove("uploads/"+image.Filename)

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusCreated,gin.H{"message":"success"})
}