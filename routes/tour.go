package routes

import (
	"github.com/forumGamers/tour-service/cmd"
	q "github.com/forumGamers/tour-service/query"
	"github.com/gin-gonic/gin"
)

func (r routes) tourRoutes(rg *gin.RouterGroup) {
	uri := rg.Group("/tour")

	uri.GET("/",q.GetAllTour)

	uri.POST("/:teamId",cmd.CreateTour)
}