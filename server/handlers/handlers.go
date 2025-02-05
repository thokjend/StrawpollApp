package handlers

import (
	"net/http"
	"strawpoll-app/database"

	"strawpoll-app/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func Register(c *gin.Context){
	var user models.User

	// Bind JSON input to the struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if user already exists
	exists, err := database.Client.Exists(database.Ctx, user.Username).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if exists > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// Hash the password before storing it
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

	// Store user in Redis
	err = database.Client.Set(database.Ctx, user.Username, hashedPassword, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account created successfully"})

}

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Get password from Redis
	storedHash, err := database.Client.Get(database.Ctx, user.Username).Result()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare hashed password with provided password
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(user.Password))
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}