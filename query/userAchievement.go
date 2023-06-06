package query

import (
	"context"
	"errors"
	"net/http"

	h "github.com/forumGamers/tour-service/helpers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserAchievement(c *gin.Context){
	user := h.GetUser(c)

	errCh := make(chan error)
	dataCh := make(chan []bson.M)

	go func(id int){
		cursor,err := getDb().Collection("userAchievement").Aggregate(context.Background(),bson.A{
			bson.D{{Key: "$match", Value: bson.D{{Key: "userId", Value: id}}}},
			bson.D{
				{Key: "$lookup",
					Value: bson.D{
						{Key: "from", Value: "achievement"},
						{Key: "localField", Value: "achievementId"},
						{Key: "foreignField", Value: "_id"},
						{Key: "as", Value: "achievement"},
					},
				},
			},
			bson.D{
				{Key: "$unwind",
					Value: bson.D{
						{Key: "path", Value: "$achievement"},
						{Key: "preserveNullAndEmptyArrays", Value: true},
					},
				},
			},
			bson.D{
				{Key: "$lookup",
					Value: bson.D{
						{Key: "from", Value: "game"},
						{Key: "localField", Value: "achievement.gameId"},
						{Key: "foreignField", Value: "_id"},
						{Key: "as", Value: "game"},
					},
				},
			},
			bson.D{
				{Key: "$unwind",
					Value: bson.D{
						{Key: "path", Value: "$game"},
						{Key: "preserveNullAndEmptyArrays", Value: true},
					},
				},
			},
		})

		if err != nil {
			if err == mongo.ErrNilCursor {
				errCh <- errors.New("Data not found")
				dataCh <- nil
				return
			}else {
				errCh <- err
				dataCh <- nil
				return
			}
		}

		var result []bson.M
		if err := cursor.All(context.Background(),&result) ; err != nil {
			errCh <- err
			dataCh <- nil
			return
		}

		errCh <- nil
		dataCh <- result
	}(user.Id)

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	data := <- dataCh

	c.JSON(http.StatusOK,data)
}