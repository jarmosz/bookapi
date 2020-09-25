package main

import (
	"book"

	"github.com/elliotforbes/go-fiber-tutorial/book"
	"github.com/gofiber/fiber"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Listen(3000)
}
