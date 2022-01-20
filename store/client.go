package store

type Client interface {
	Set(shortUrl string, originalUrl string)
	Get(shortUrl string) string
}
