package main

import (
	"github.com/fuboki10/nanoURL/handler"
	"github.com/fuboki10/nanoURL/service/store"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	store.InitializeStore();

	initRoutes(app);

	startServer(app);
}

func initRoutes(app *fiber.App) {
	app.Static("/","./public")

	app.Post("/urls",  handler.CreateUrl)
	app.Get("/urls/:url",  handler.UrlRedirect)
}

func startServer(app *fiber.App)  {
	err := app.Listen(":3000")
	if err != nil {
		panic("failed to start server")
	}
}

