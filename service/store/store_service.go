package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
    ctx = context.Background()
)


const CacheDuration = 6 * time.Hour


// Initializing the store service and return a store pointer 
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully")
	storeService.redisClient = redisClient
	return storeService
}


func SaveUrl(shortUrl string, originalUrl string) {
	storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
}

func RetrieveInitialUrl(shortUrl string) string {
	result, _ := storeService.redisClient.Get(ctx, shortUrl).Result()
	return result
}