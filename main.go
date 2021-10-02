package main

import (
	"os"

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
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := app.Listen(port)
	if err != nil {
		panic("failed to start server")
	}
}

