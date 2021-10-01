package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://github.com/Hassan950/OudBackEnd/blob/master/src/middlewares/rateLimiter.js"
	shortURL := "Jsz4k57oAX"

	SaveUrl(shortURL, initialLink)

	assert.Equal(t, initialLink, RetrieveInitialUrl(shortURL))
}