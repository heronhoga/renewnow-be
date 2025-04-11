package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CheckSession(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	fmt.Println(token)

	return ctx.Next()
}