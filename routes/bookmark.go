package routes

import (
	"github.com/forumGamers/tour-service/cmd"
	md "github.com/forumGamers/tour-service/middlewares"
	"github.com/gin-gonic/gin"
)

func (r routes) bookmarkRoutes(rg *gin.RouterGroup){
	uri := rg.Group("/bookmark")

	uri.POST("/:tourId", md.Authentication, cmd.AddBookmark)
}