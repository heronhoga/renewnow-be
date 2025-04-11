package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/heronhoga/renewnow-be/config"
	"github.com/heronhoga/renewnow-be/models"
	"github.com/heronhoga/renewnow-be/requests"
)

func CreateLicense(ctx *fiber.Ctx) error {
 	newLicense := new(requests.CreateLicenseRequest)

	if err := ctx.BodyParser(newLicense); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "something wrong with your request",
			"error": err,
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

	 //find owner
	 var existingUser models.User

	 authToken := ctx.Get("Authorization")[7:]

	 findToken := config.DB.Where("session = ?", authToken).First(&existingUser)
	 if findToken.RowsAffected == 0 {
		 return ctx.Status(403).JSON(fiber.Map{
			 "message": "forbidden",
		 })
	 }

	 //map request to model
	 newLicenseData := models.License{
		LicenseType: newLicense.LicenseType,
		Expire: newLicense.Expire,
		UserID: existingUser.ID ,
	 }

	 errCreateLicense := config.DB.Create(&newLicenseData).Error

	 if errCreateLicense != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	 }

	 return ctx.Status(200).JSON(fiber.Map{
		"Message": "License successfully created",
		"License": newLicenseData,
	 })
}