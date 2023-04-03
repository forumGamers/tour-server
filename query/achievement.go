package query

import (
	"context"
	"errors"
	"net/http"

	l "github.com/forumGamers/tour-service/loaders"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getDb() *mongo.Database {
	return l.GetDb()
}

func GetByGameId(c *gin.Context){
	gameId := c.Param("gameId")

	id,err := primitive.ObjectIDFromHex(gameId)

	if err != nil {
		panic("Invalid ObjectID")
	}

	dataCh := make(chan []bson.M)
	errCh := make(chan error)

	go func(id primitive.ObjectID){
		cursor,err := getDb().Collection("achievement").Find(context.Background(),bson.M{"gameId":id})
	
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

		errCh <- nil
		dataCh <- data
	}(id)

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	data := <- dataCh

	c.JSON(http.StatusOK,data)
}