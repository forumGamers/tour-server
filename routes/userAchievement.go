package routes

import (
	"github.com/forumGamers/tour-service/cmd"
	md "github.com/forumGamers/tour-service/middlewares"
	"github.com/gin-gonic/gin"
)

func (r routes) UserAchievementRoutes(rg *gin.RouterGroup){
	uri := rg.Group("/user-achievement")

	uri.POST("/:id",md.Authentication,cmd.CreateUserAchievement)
}