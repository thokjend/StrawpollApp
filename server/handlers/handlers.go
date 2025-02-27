package handlers

import (
	"net/http"
	"strawpoll-app/database"
	"strawpoll-app/websocket"
	"strings"
	"time"

	"strawpoll-app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
)

/*
gin.Context = 	pointer to a Context object from the Gin web framework,
				which is used to handle HTTP requests and responses.
				It provides methods to parse JSON, set response status codes, and send JSON responses.
*/

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

	// Create a session token
	sessionToken := uuid.NewString()
	expiresAt := 30 * time.Minute
	key := "session:" + sessionToken 

	// Store session token in Redis with expiration time
	err = database.Client.SetEx(database.Ctx, key, user.Username, expiresAt).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store session token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   sessionToken,
	})
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Look up the token in Redis
        username, err := database.Client.Get(database.Ctx, "session:"+token).Result()
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session token"})
            c.Abort()
            return
        }

		err = database.Client.Expire(database.Ctx, "session:"+token, 30*time.Minute).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not refresh session"})
			c.Abort()
			return
		}


        c.Set("username", username) // Store username in context
        c.Next()
    }
}


func CreatePoll(c *gin.Context) {
	var pollData models.Poll
	id, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 8)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate poll ID"})
		return
	}

	if err := c.ShouldBindJSON(&pollData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Store poll metadata
	err = database.Client.HSet(database.Ctx, id, map[string]interface{}{
		"title":        pollData.Title,
		"description":  pollData.Description,
		"multiple":     pollData.Settings[0],
		"requireNames": pollData.Settings[1],
	}).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save poll data"})
		return
	}

	// Store options with an initial vote count of 0
	optionVotes := make(map[string]interface{})
	for _, option := range pollData.Options {
		optionVotes[option] = 0
	}

	err = database.Client.HSet(database.Ctx, id+":votes", optionVotes).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save poll options"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Poll created successfully",
		"id":      id,
	})
}


func GetPoll(c *gin.Context) {
	pollID := c.Param("id")

	poll, err := database.Client.HGetAll(database.Ctx, pollID).Result()
	if err != nil || len(poll) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
		return
	}

	// Retrieve vote counts
	votes, err := database.Client.HGetAll(database.Ctx, pollID+":votes").Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve votes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Poll retrieved successfully",
		"data": gin.H{
			"title":        poll["title"],
			"description":  poll["description"],
			"multiple":     poll["multiple"],
			"requireNames": poll["requireNames"],
			"votes":        votes,
		},
	})
}


func VotePoll(c *gin.Context) {
	pollID := c.Param("id")
	var voteData struct {
		Option string `json:"Option"`
	}

	if err := c.ShouldBindJSON(&voteData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	options := strings.Split(voteData.Option, ",")

	// Increment the vote count for the chosen option
	for _, option := range options {
		err := database.Client.HIncrBy(database.Ctx, pollID+":votes", option, 1).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record vote"})
			return
		}
	}

	// Sends a WebSocket message to all connected clients
	websocket.BroadcastMessage("update")
	// Any client listening to the WebSocket will recive "update", which can be used to refresh the poll results in real time

	c.JSON(http.StatusOK, gin.H{"message": "Vote recorded successfully"})
}