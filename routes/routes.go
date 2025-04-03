package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heronhoga/renewnow-be/controllers"
)

func Route(r *fiber.App) {
	r.Post("/register", controllers.Register)
}