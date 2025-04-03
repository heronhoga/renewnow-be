package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/heronhoga/renewnow-be/config"
	"github.com/heronhoga/renewnow-be/models"
	"github.com/heronhoga/renewnow-be/requests"
	"github.com/heronhoga/renewnow-be/utils"
)

func Register(ctx *fiber.Ctx) error {
	//get request data
	user := new(requests.CreateUserRequest)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "something wrong with your request",
		})
	}

	//validate requests
	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "your request data is not valid",
		})
	}

	var existingUser models.User

	//email check result
	emailCheckResult := config.DB.Where("email = ?", user.Email).First(&existingUser)
	if emailCheckResult.RowsAffected > 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	//username check result
	usernameCheckResult := config.DB.Where("username = ?", user.Username).First(&existingUser)
	if usernameCheckResult.RowsAffected > 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	//hash password
	newUserHashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	//map request to model
	newUser := models.User{
		Username: user.Username,
		Password: newUserHashedPassword,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
	}

	//create new user data
	errCreateUser := config.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
			"error":   errCreateUser,
		})
	}

	//return success response
	return ctx.JSON(fiber.Map{
		"message": "user successfully created",
	})

}