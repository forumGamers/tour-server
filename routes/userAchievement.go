package routes

import (
	"github.com/forumGamers/tour-service/cmd"
	md "github.com/forumGamers/tour-service/middlewares"
	q "github.com/forumGamers/tour-service/query"
	"github.com/gin-gonic/gin"
)

func (r routes) UserAchievementRoutes(rg *gin.RouterGroup){
	uri := rg.Group("/user-achievement")

	uri.POST("/:id",md.Authentication,cmd.CreateUserAchievement)

	uri.GET("/",md.Authentication,q.GetUserAchievement)
}