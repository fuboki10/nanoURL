package store

import (
	"fmt"
	"time"

	"github.com/fuboki10/nanoURL/store"
	"github.com/fuboki10/nanoURL/store/fake"
	"github.com/fuboki10/nanoURL/store/redis"
)

// Define the struct wrapper around raw Redis client
type StorageService struct {
	client store.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
)


const CacheDuration = 6 * time.Hour


// Initializing the store service and return a store pointer 
func InitializeStore() *StorageService {
	redisClient, err := redis.New()

	if err != nil {
		// connect to temp clinet
		tempRedis, _ := fake.New()
		storeService.client = tempRedis
	} else {
		fmt.Printf("\nRedis started successfully")
		storeService.client = redisClient
	}

	return storeService
}


func SaveUrl(shortUrl string, originalUrl string) {
	storeService.client.Set(shortUrl, originalUrl)
}

func RetrieveInitialUrl(shortUrl string) string {
	result := storeService.client.Get(shortUrl)
	return result
}