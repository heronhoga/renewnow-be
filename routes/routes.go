package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heronhoga/renewnow-be/controllers"
	"github.com/heronhoga/renewnow-be/middlewares"
)

func Route(r *fiber.App) {
	r.Post("/register", controllers.Register)
	r.Post("/login", controllers.Login)
	r.Post("/logout", middlewares.CheckSession, controllers.Logout)
}