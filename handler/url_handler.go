package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/fuboki10/nanoURL/service"
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

	shortUrl := service.shortener.GenerateShortLink(creationRequest.LongUrl)
	service.store.SaveUrl(shortUrl, creationRequest.LongUrl)

	host := "http://localhost:3000/"
	url := host + shortUrl

	return ctx.SendString(url)
}

func HandleShortUrlRedirect(ctx *fiber.Ctx) {
	shortUrl := ctx.Params("url")
	url := service.store.RetrieveInitialUrl(shortUrl)
	ctx.Redirect(url)
}