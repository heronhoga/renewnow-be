package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heronhoga/renewnow-be/config"
	"github.com/heronhoga/renewnow-be/models"
)

func CheckSession(ctx *fiber.Ctx) error {
	var user models.User
	token := ctx.Get("Authorization")

	if token == "" {
		return ctx.Status(403).JSON(fiber.Map{
			"message": "forbidden",
		})
	}
	
	token = token[7:]

	//find matching session token
	findToken := config.DB.Where("session = ?", token).First(&user)
	if findToken.RowsAffected == 0 {
		return ctx.Status(403).JSON(fiber.Map{
			"message": "forbidden",
		})
	}
	

	return ctx.Next()
}