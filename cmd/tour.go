package cmd

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	h "github.com/forumGamers/tour-service/helpers"
	m "github.com/forumGamers/tour-service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTour(c *gin.Context){
	host := c.Param("teamId")

	name,gameId,pricePool,slot,startDate,fee,location,description,tags :=
	c.PostForm("name"),
	c.PostForm("gameId"),
	c.PostForm("pricePool"),
	c.PostForm("slot"),
	c.PostForm("startDate"),
	c.PostForm("registrationFee"),
	c.PostForm("location"),
	c.PostForm("description"),
	c.PostForm("tags")

	_,err := primitive.ObjectIDFromHex(host)

	if err != nil {
		panic("Invalid ObjectID")
	}

	id,err := primitive.ObjectIDFromHex(gameId)

	if err != nil {
		panic("Invalid ObjectID")
	}

	errCh := make(chan error)

	go func(
		host string,
		name string,
		gameId primitive.ObjectID,
		pricePool string,
		slots string,
		startDate string,
		fees string,
		location string,
		description string,
		tags string,
	){
		pool,err := strconv.ParseInt(pricePool,10,64)

		if err != nil {
			errCh <- errors.New("Invalid data")
			return
		}

		slot,err := strconv.ParseInt(slots,10,64)

		if err != nil {
			errCh <- errors.New("Invalid data")
			return
		}

		date,err := time.Parse("2023-12-30",startDate)

		if err != nil {
			errCh <- errors.New("Invalid data")
			return
		}

		fee,err := strconv.ParseInt(fees,10,64)

		if err != nil {
			errCh <- errors.New("Invalid data")
			return
		}

		if err := h.ValidateInput(map[string]string{
			"name":name,
			"location":location,
			"description":description,
		}) ; err != nil {
			errCh <- err
			return
		}

		tag := strings.Split(tags,",")

		if _,err := getDb().Collection("tour").InsertOne(context.Background(),m.Tour{
			Host: host,
			Name: name,
			GameId: gameId,
			PricePool: int(pool),
			Slot: int(slot),
			StartDate: date,
			Location: location,
			RegistrationFee: int(fee),
			Description: description,
			Tags: tag,
		}) ; err != nil {
			errCh <- err
			return
		}
	}(
		host,
		name,
		id,
		pricePool,
		slot,
		startDate,
		fee,
		location,
		description,
		tags,
	)
}