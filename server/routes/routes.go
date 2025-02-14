package routes

import (
	"strawpoll-app/handlers"
	"time"

	"strawpoll-app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/polls", handlers.CreatePoll)
		protected.GET("/polls/:id", handlers.GetPoll)
		protected.POST("/polls/:id/vote", handlers.VotePoll)
	}
	
	return router
}