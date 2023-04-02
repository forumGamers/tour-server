package routes

import (
	"github.com/forumGamers/tour-service/cmd"
	q "github.com/forumGamers/tour-service/query"
	"github.com/gin-gonic/gin"
)

func (r routes) gameRoutes(rg *gin.RouterGroup){
	uri := rg.Group("/game")

	uri.POST("/",cmd.CreateGame)

	uri.GET("/",q.GetAllGame)
}