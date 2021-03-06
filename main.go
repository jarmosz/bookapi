package main

import (
	"github.com/jarmosz/books"

	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", books.GetBooks)
	app.Get("/api/v1/book/:id", books.GetBook)
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Listen(3000)
}
