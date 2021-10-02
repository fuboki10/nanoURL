package redis

import (
	"time"

	"github.com/go-redis/redis"
)


type RedisClient struct {
	redisClient *redis.Client
}

const CacheDuration = 6 * time.Hour

func New() (*RedisClient, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := redisClient.Ping().Result()

	client := &RedisClient{redisClient : redisClient}

	return client, err
}

func (client *RedisClient) Set(shortUrl string, originalUrl string) {
	client.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
}
	
func (client *RedisClient) Get(shortUrl string) string {
	result, _ := client.redisClient.Get(shortUrl).Result()
	return result;
}