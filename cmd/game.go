package cmd

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	cfg "github.com/forumGamers/tour-service/config"
	h "github.com/forumGamers/tour-service/helpers"
	"github.com/forumGamers/tour-service/loaders"
	m "github.com/forumGamers/tour-service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func getDb()*mongo.Database {
	return loaders.GetDb()
}

func CreateGame(c *gin.Context){
	
	name,types,description := c.PostForm("name"),c.PostForm("type"),c.PostForm("description")

	if name == "" || types == "" || description == "" {
		panic("Invalid data")
	}

	image,err := c.FormFile("image")

	if err != nil {
		panic(err.Error())
	}

	if err := h.ValidateInput(map[string]string{
		"name":name,
		"types":types,
		"description":description,
	}) ; err != nil {
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
	fileidCh := make(chan string)
	errUpload := make(chan error)
	errCh := make(chan error)

	go func(data []byte,imageName string){
		url,fileId,err := cfg.UploadImage(data,imageName,"gameImage")

		if err != nil {
			errUpload <- errors.New("Bad Gateway")
			urlCh <- ""
			fileidCh <- ""
			return
		}

		errUpload <- nil
		urlCh <- url
		fileidCh <- fileId

	}(data,image.Filename)

	go func(name string,types string,description string){
		if err := <- errUpload ; err != nil {
			errCh <- err
			return
		}


		if _,err := getDb().Collection("game").InsertOne(context.Background(),m.Game{
			Image: <- urlCh,
			ImageId: <- fileidCh,
			Name: name,
			Type: types,
			Description: description,
		}) ; err != nil {
			errCh <- err
			return
		}

		errCh <- nil
	}(name,types,description)

	file.Close()

	os.Remove("uploads/"+image.Filename) 

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusCreated,gin.H{"message":"success"})
}