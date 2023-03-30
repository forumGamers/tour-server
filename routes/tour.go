package routes

import "github.com/gin-gonic/gin"

func (r routes) tourRoutes(rg *gin.RouterGroup) {
	uri := rg.Group("/tour")

	uri.GET("/")
}