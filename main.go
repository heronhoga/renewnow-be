package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heronhoga/renewnow-be/config"
	"github.com/heronhoga/renewnow-be/models"
	"github.com/heronhoga/renewnow-be/routes"
)

func main() {
	app := fiber.New()
	config.ConnectDb()
	
	//migration
	config.MigrateDB(&models.User{}, &models.License{})

	routes.Route(app)
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}