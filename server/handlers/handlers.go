package handlers

import (
	"fmt"
	"net/http"
	"strawpoll-app/database"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	val, err := database.Client.Get(database.Ctx, "name").Result()
	if err != nil {
		// Respond with an error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting key"})
		fmt.Println("Error getting key:", err)
		return
	}

	// Respond with the retrieved value
	c.JSON(http.StatusOK, gin.H{"value": val})
}