package query

import (
	"context"
	"errors"
	"net/http"

	h "github.com/forumGamers/tour-service/helpers"
	m "github.com/forumGamers/tour-service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMyTeam(c *gin.Context){
	user := h.GetUser(c)

	errCh := make(chan error)
	dataCh := make(chan []bson.M)

	go func (user m.User)  {
		cursor,err := getDb().Collection("team").Find(context.Background(),bson.M{"ownerId":user.Id})

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

		defer cursor.Close(context.Background())

		var data []bson.M
		for cursor.Next(context.Background()){
			var result bson.M
			if err := cursor.Decode(&result) ; err != nil {
				errCh <- err
				dataCh <- nil
				return
			}
			data = append(data,result)
		}
		
		if len(data) < 1 {
			errCh <- errors.New("Data not found")
			dataCh <- nil
			return
		}

		errCh <- nil
		dataCh <- data
	}(user)

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	data := <- dataCh

	c.JSON(http.StatusOK,data)
}

func GetTeamById(c *gin.Context){
	teamId := c.Param("teamId")

	id,err := primitive.ObjectIDFromHex(teamId)

	if err != nil {
		panic("Invalid ObjectID")
	}

	errCh := make(chan error)
	dataCh := make(chan []bson.M)

	go func(id primitive.ObjectID){
		cursor,err := getDb().Collection("teamAchievement").Aggregate(context.Background(),bson.A{
			bson.D{{Key: "$match", Value: bson.D{{Key: "teamId", Value: id}}}},
			bson.D{
				{Key: "$lookup",
					Value: bson.D{
						{Key: "from", Value: "team"},
						{Key: "localField", Value: "teamId"},
						{Key: "foreignField", Value: "_id"},
						{Key: "as", Value: "team"},
					},
				},
			},
			bson.D{
				{Key: "$unwind",
					Value: bson.D{
						{Key: "path", Value: "$team"},
						{Key: "preserveNullAndEmptyArrays", Value: true},
					},
				},
			},
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
	}(id)

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	data := <- dataCh

	c.JSON(http.StatusOK,data)
}
