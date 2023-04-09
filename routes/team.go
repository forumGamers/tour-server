package routes

import (
	"github.com/forumGamers/tour-service/cmd"
	md "github.com/forumGamers/tour-service/middlewares"
	q "github.com/forumGamers/tour-service/query"
	"github.com/gin-gonic/gin"
)

func (r routes) teamRoutes(rg *gin.RouterGroup){
	uri := rg.Group("/team")

	uri.POST("/", md.Authentication, cmd.CreateTeam)

	uri.GET("/myTeam",md.Authentication,q.GetMyTeam)

	uri.GET("/:teamId",q.GetTeamById)
}