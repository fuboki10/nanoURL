package handler

import (
	"bytes"

	shortener "github.com/fuboki10/nanoURL/service/shortner"
	"github.com/fuboki10/nanoURL/service/store"
	"github.com/gofiber/fiber/v2"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func CreateUrl(ctx *fiber.Ctx) error {
	var creationRequest UrlCreationRequest
	err := ctx.BodyParser(&creationRequest); 
	if err != nil {
		return ctx.SendString("ERROR")
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
	store.SaveUrl(shortUrl, creationRequest.LongUrl)

	var buffer bytes.Buffer

	buffer.WriteString(ctx.Hostname())
	buffer.WriteString("/")
	buffer.WriteString(shortUrl)

	url := buffer.String()

	return ctx.SendString(url)
}

func UrlRedirect(ctx *fiber.Ctx) error {
	shortUrl := ctx.Params("url")
	url := store.RetrieveInitialUrl(shortUrl)
	return ctx.Redirect(url)
}