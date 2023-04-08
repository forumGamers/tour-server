package cmd

import (
	"context"
	"net/http"
	"time"

	h "github.com/forumGamers/tour-service/helpers"
	md "github.com/forumGamers/tour-service/middlewares"
	m "github.com/forumGamers/tour-service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddBookmark(c *gin.Context){

	user := h.GetUser(c)

	tourId := c.Param("tourId")

	id,err := primitive.ObjectIDFromHex(tourId)

	if err != nil {
		panic("Invalid ObjectID")
	}

	errCh := make(chan error)

	go func (user md.User,id primitive.ObjectID)  {
		if _,err := getDb().Collection("bookmark").InsertOne(context.Background(),m.Bookmark{
			UserId: user.Id,
			TourId: id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}) ; err != nil {
			errCh <- err
			return
		}
	}(user,id)

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusCreated,gin.H{"message":"success"})
}