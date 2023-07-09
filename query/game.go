package query

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllGame(c *gin.Context){

	errCh := make(chan error)
	dataCh := make(chan []bson.M)

	go func(){
		cursor,err := getDb().Collection("game").Find(context.Background(),bson.M{},options.Find().SetSort(bson.D{{Key: "createdAt",Value: -1}}))

		if err != nil {
			errCh <- err
			dataCh <- nil
			return
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
			data = append(data, result)
		}
		
		if err := cursor.Err() ; err != nil {
			errCh <- err
			dataCh <- nil
			return
		}

		if len(data) < 1 {
			errCh <- errors.New("Data not found")
			dataCh <- nil
			return
		}

		errCh <- nil
		dataCh <- data
	}()

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	data := <- dataCh

	c.JSON(http.StatusOK,data)
}