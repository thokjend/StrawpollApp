package routes

import (
	"net/http"
	"strawpoll-app/handlers"
	"strawpoll-app/websocket"
	"time"

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
	router.POST("/create", handlers.CreatePoll)
	router.POST("/poll/:id/vote", handlers.VotePoll)

	router.GET("/poll/:id", handlers.GetPoll)
	router.GET("/polls", handlers.GetPolls)
	router.GET("/ws", websocket.HandleWebSocket)
	router.GET("/protected-route", handlers.AuthMiddleware(), func(c *gin.Context) {
		username, _ := c.Get("username")
		c.JSON(http.StatusOK, gin.H{"message": "Access granted", "username": username})
	})
	
	return router
}