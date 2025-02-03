package main

import (
	"strawpoll-app/database"
	"strawpoll-app/routes"
)

func main() {
	database.InitRedis()
    router := routes.SetupRoutes()
    
    router.Run("localhost:8080")
}