package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("public/index.html")
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		return controller.createUrl(ctx)
	})

	app.Get("/:url", func(ctx *fiber.Ctx) error {
		return controller.redirect(ctx)
	})

	err := app.Listen(":3000")
	if err != nil {
		panic("failed to start server")
	}
}