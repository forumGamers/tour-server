package cmd

import (
	"context"
	"errors"
	"net/http"
	"time"

	h "github.com/forumGamers/tour-service/helpers"
	m "github.com/forumGamers/tour-service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserAchievement(c *gin.Context){
	user := h.GetUser(c)

	id := c.Param("id")

	Id,err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic("Invalid ObjectID")
	}

	errCheck := make(chan error)
	errCh := make(chan error)

	go func (id primitive.ObjectID)  {
		if err := getDb().Collection("achievement").FindOne(context.Background(),bson.M{"_id":id}).Err() ; err != nil {
			if err == mongo.ErrNilDocument {
				errCheck <- errors.New("Data not found")
				return
			}else {
				errCheck <- err
				return
			}
		}

		errCheck <- nil
	}(Id)

	go func (id int,achievementId primitive.ObjectID)  {
		if err := <- errCheck ; err != nil {
			errCh <- err
			return
		} 

		if _,err := getDb().Collection("userAchievement").InsertOne(context.Background(),m.UserAchievement{
			AchievementId: achievementId,
			UserId: id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}) ; err != nil {
			errCh <- err
			return
		}

		errCh <- nil
	}(user.Id,Id)

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusCreated,gin.H{"message":"success"})
}