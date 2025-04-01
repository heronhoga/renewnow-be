package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heronhoga/renewnow-be/config"
)

func main() {
	app := fiber.New()
	config.ConnectDb()
	// run this function once
	// config.MigrateDB(&models.User{}, &models.License{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}