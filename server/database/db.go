package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx    = context.Background()
	Client *redis.Client
)

func InitRedis(){
	// Connect to Memurai (redis)
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Memurai's default address
		DB:   0,                // Use default DB
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Memurai:", err)
		return
	}
	fmt.Println("Connected to Memurai!")
}