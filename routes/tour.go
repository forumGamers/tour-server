package routes

import (
	"github.com/forumGamers/tour-service/cmd"
	"github.com/gin-gonic/gin"
)

func (r routes) tourRoutes(rg *gin.RouterGroup) {
	uri := rg.Group("/tour")

	uri.GET("/")

	uri.POST("/:teamId",cmd.CreateTour)
}