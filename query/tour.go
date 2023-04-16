package query

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTour(c *gin.Context){
	gameId,minDate,maxDate,page,limit := 
	c.Query("gameId"),
	c.Query("minDate"),
	c.Query("maxDate"),
	c.Query("page"),
	c.Query("limit")

	errCh := make(chan error)
	dataCh := make(chan []bson.M)

	go func(
		gameId string,
		minDate string,
		maxDate string,
		page string,
		limit string,
	){
		pipeline := bson.A{}
		matchStage := bson.M{ "$match":bson.M{} }
		var pg int
		var lmt int

		if gameId != "" {
			id,err := primitive.ObjectIDFromHex(gameId)

			if err != nil {
				errCh <- errors.New("Invalid ObjectID")
				dataCh <- nil
				return
			}

			matchStage["$match"].(bson.M)["gameId"] = id
		}

		if minDate != "" || maxDate != "" {
			dateFilter := bson.M{}

			if minDate != "" {
				min,err := time.Parse("02-01-2006",minDate)

				if err != nil {
					errCh <- err
					dataCh <- nil
					return
				}

				dateFilter["$gte"] = min
			}

			if maxDate != "" {
				max,err := time.Parse("02-01-2006",maxDate)

				if err != nil {
					errCh <- err
					dataCh <- nil
					return
				}

				dateFilter["$lte"] = max
			}

			matchStage["$match"].(bson.M)["createdAt"] = dateFilter
		}

		if len(matchStage) > 0 {
			pipeline = append(pipeline, matchStage)
		}

		if page != ""  {
			p,err := strconv.Atoi(page)

			if err != nil {
				errCh <- errors.New("Invalid params")
				dataCh <- nil
				return
			}

			pg = p
		}else {
			pg = 1
		}

		if limit != "" {
			l,err := strconv.Atoi(limit)

			if err != nil {
				errCh <- errors.New("Invalid params")
				dataCh <- nil
				return
			}

			lmt = l
		}else {
			lmt = 10
		}

		pipeline = append(pipeline, bson.M{"$skip":(pg - 1) * lmt})
		pipeline = append(pipeline, bson.M{"$limit":lmt})

		cursor,err := getDb().Collection("tour").Aggregate(context.Background(),pipeline)

		if err != nil {
			errCh <- err
			dataCh <- nil
			return
		}

		var results []bson.M
		for cursor.Next(context.Background()) {
			var result bson.M
			if err := cursor.Decode(&result); err != nil {
				errCh <- err
				dataCh <- nil
				return
			}
			results = append(results, result)
		}

		if len(results) < 1 {
			errCh <- errors.New("Data not found")
			dataCh <- nil
			return
		}

		errCh <- nil
		dataCh <- results
	}(
		gameId,
		minDate,
		maxDate,
		page,
		limit,
	)

	if err := <- errCh ; err != nil {
		panic(err.Error())
	}

	data := <- dataCh

	c.JSON(http.StatusOK,data)
}