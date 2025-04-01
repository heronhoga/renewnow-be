package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heronhoga/renewnow-be/config"
)

func main() {
	app := fiber.New()
	config.ConnectDb()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}