package routes

import (
	"net/http"
	"os"

	md "github.com/forumGamers/tour-service/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type routes struct {
	router *gin.Engine
}

func Routes(){
	if err := godotenv.Load() ; err != nil {
		panic(err.Error())
	}

	r := routes { router: gin.Default() }

	c := cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("CORSLIST")},
		AllowMethods: []string{"GET","POST","PUT","DELETE","PATCH","OPTIONS"},
		AllowHeaders: []string{"Content-Type","Authorization"},
		AllowCredentials: true,
	})

	r.router.Use(func (c *gin.Context){
		if c.Request.Method != "OPTIONS" {
			origin := c.Request.Header.Get("Origin")
			if origin == "" {
                c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
                return
            }
		}
		c.Next()
	})

	r.router.Use(c)

	r.router.Use(logger.SetLogger())

	r.router.Use(md.ErrorHandler)

	//testing connection
	r.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"pong"})
	})

	groupRoutes := r.router.Group("/api")

	r.tourRoutes(groupRoutes)

	r.gameRoutes(groupRoutes)

	r.achievementRoutes(groupRoutes)

	r.bookmarkRoutes(groupRoutes)

	r.teamRoutes(groupRoutes)

	port := os.Getenv("PORT")

	if port == "" {
		port = "4200"
	}

	r.router.Run(":"+port)
}