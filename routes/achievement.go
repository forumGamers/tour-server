package routes

import (
	"github.com/forumGamers/tour-service/cmd"
	"github.com/gin-gonic/gin"
)

func (r routes) achievementRoutes(rg *gin.RouterGroup){
	uri := rg.Group("/achievement")

	uri.POST("/:gameId",cmd.CreateAchievement)
}