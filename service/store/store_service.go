package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
	tempClinet map[string]string
}

var tempRedis map[string]string

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
)


const CacheDuration = 6 * time.Hour


// Initializing the store service and return a store pointer 
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		// connect to temp clinet
		tempRedis = make(map[string]string)
		storeService.tempClinet = tempRedis
	} else {
		fmt.Printf("\nRedis started successfully")
		storeService.redisClient = redisClient
	}

	return storeService
}


func SaveUrl(shortUrl string, originalUrl string) {
	if storeService.redisClient == nil {
		storeService.tempClinet[shortUrl] = originalUrl
	} else {
		storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	result := ""
	if storeService.redisClient == nil {
		result = storeService.tempClinet[shortUrl]
	} else {
		result, _ = storeService.redisClient.Get(shortUrl).Result()
	}
	return result
}