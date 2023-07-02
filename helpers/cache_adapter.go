package helpers

import (
	"time"

	"github.com/cosmos-sajal/go_boilerplate/initializers"
	"github.com/go-redis/redis"
)

func getRedisClient() *redis.Client {
	return initializers.RedisClient
}

func SetCacheValue(key string, value string, expiry ...int) error {
	var timeout time.Duration = 0
	if len(expiry) > 0 {
		timeout = time.Duration(expiry[0]) * time.Second
	}

	redisClient := getRedisClient()
	_, err := redisClient.Set(key, value, timeout).Result()
	if err != nil {
		return err
	}

	return nil
}

func GetCacheValue(key string) (string, error) {
	redisClient := getRedisClient()
	val, err := redisClient.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func DeleteKey(key string) error {
	redisClient := getRedisClient()
	_, err := redisClient.Del(key).Result()
	if err != nil {
		return err
	}

	return nil
}
