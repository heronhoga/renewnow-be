package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/heronhoga/renewnow-be/config"
	"github.com/heronhoga/renewnow-be/models"
	"github.com/heronhoga/renewnow-be/requests"
	"github.com/heronhoga/renewnow-be/utils"
	"golang.org/x/crypto/bcrypt"

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

	//phone check result
	phoneCheckResult := config.DB.Where("phone = ?", user.Phone).First(&existingUser)
	if phoneCheckResult.RowsAffected > 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Phone number already exists",
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
		Phone: user.Phone,
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

func Login(ctx *fiber.Ctx) error {
	user := new(requests.LoginUserRequest)

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

	//username and password validation
	var existingUser models.User
	usernameCheckResult := config.DB.Where("username = ?", user.Username).First(&existingUser)
	
	if usernameCheckResult.RowsAffected == 0 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Username and/or password is wrong",
		})
	}
	
	errCompare := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if errCompare != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Username and/or password is wrong",
		})
	}
	
	//generate token
	sessionToken, errGenerate := utils.GenerateToken();
	if errGenerate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	//update session column
	updateToken := config.DB.Model(&models.User{}).Where("username = ?", user.Username).Update("session", sessionToken)
	if updateToken.RowsAffected == 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Login successful",
		"token": sessionToken,
		"username": user.Username,
	})
}

func Logout(ctx *fiber.Ctx) error {
	userLogout := new(requests.LogoutUserRequest)

	if err := ctx.BodyParser(userLogout); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "something wrong with your request",
		})
	}

	//set session column to null
	updateToken := config.DB.Model(&models.User{}).Where("session = ?", userLogout.Session).Update("session", nil)
	if updateToken.RowsAffected == 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Logout successful",
	})
}