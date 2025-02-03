package routes

import (
	"strawpoll-app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/test", handlers.Test)
	//router.POST("/register")
	
	return router
}