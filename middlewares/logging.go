package middlewares

import (
	"context"
	"fmt"
	"time"

	h "github.com/forumGamers/tour-service/helpers"
	"github.com/forumGamers/tour-service/loaders"
	m "github.com/forumGamers/tour-service/models"
	"github.com/gin-gonic/gin"
)

func Logging(c *gin.Context){
	defer func(){
		id := h.GetUser(c).Id

		responseTime := c.MustGet("start").(time.Time)

		if _,err := loaders.GetDb().Collection("log").InsertOne(context.Background(),m.Log{
			Path: c.Request.URL.Path,
			UserId: id ,
			Method: c.Request.Method,
			StatusCode: c.Writer.Status(),
			Origin: c.Request.Header.Get("Origin"),
			ResponseTime: int(time.Since(responseTime).Milliseconds()),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}) ; err != nil {
			fmt.Println(err)
			return
		}
	}()

	c.Next()
}