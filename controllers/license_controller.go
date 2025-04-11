package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/heronhoga/renewnow-be/requests"
)

func CreateLicense(ctx *fiber.Ctx) error {
 	newLicense := new(requests.CreateLicenseRequest)

	if err := ctx.BodyParser(newLicense); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "something wrong with your request",
		})
	}

	//validate requests
	 validate := validator.New()

	 errValidate := validate.Struct(newLicense)

	 if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "your request data is not valid",
		})
	 }

	 //not done
	 return errValidate
}