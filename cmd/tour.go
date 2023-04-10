package cmd

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	h "github.com/forumGamers/tour-service/helpers"
	m "github.com/forumGamers/tour-service/models"
	v "github.com/forumGamers/tour-service/validation"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTour(c *gin.Context){
	host := c.Param("teamId")

	name,gameId,pricePool,slots,startDate,registrationFee,location,description,tags :=
	c.PostForm("name"),
	c.PostForm("gameId"),
	c.PostForm("pricePool"),
	c.PostForm("slot"),
	c.PostForm("startDate"),
	c.PostForm("registrationFee"),
	c.PostForm("location"),
	c.PostForm("description"),
	c.PostForm("tags")

	id,pool,slot,date,fee,err := v.ValidateCreateTour(
		host,
		name,
		gameId,
		pricePool,
		slots,
		startDate,
		registrationFee,
		location,
		description,
		tags,
	)

	if err != nil {
		panic(err.Error())
	}

	image,err := c.FormFile("image")

	if err != nil {
		panic("Invalid data")
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

	go h.UploadImage(data,image.Filename,"tourImage",urlCh,fileIdCh,errUpload)

	go func(
		host string,
		name string,
		gameId primitive.ObjectID,
		pricePool int,
		slot int,
		date time.Time,
		fee int,
		location string,
		description string,
		tags string,
	){
		if err := <- errUpload ; err != nil {
			errCh <- err
			return
		}

		tag := strings.Split(tags,",")

		if _,err := getDb().Collection("tour").InsertOne(context.Background(),m.Tour{
			Host: host,
			Name: name,
			Image: <- urlCh,
			ImageId: <- fileIdCh,
			GameId: gameId,
			PricePool: int(pool),
			Slot: int(slot),
			StartDate: date,
			Location: location,
			Status: "Preparation",
			RegistrationFee: int(fee),
			Description: description,
			Tags: tag,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}) ; err != nil {
			errCh <- err
			return
		}

		errCh <- nil
	}(
		host,
		name,
		id,
		pool,
		slot,
		date,
		fee,
		location,
		description,
		tags,
	)

	file.Close()

	os.Remove("uploads/"+image.Filename)

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusCreated,gin.H{"message":"success"})
}