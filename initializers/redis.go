package initializers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func ConnectToRedis() {
	fmt.Println("Testing Golang Redis")
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       redisDB,
	})

	// Ping the Redis server to test the connection
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)
}
