package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	url := "https://github.com/fuboki10/Ketofan-Back-End/blob/main/src/routes/v1/auth.router.ts"
	shortUrl := GenerateShortLink(url)

	assert.Equal(t, "emFeKWge", shortUrl)
}