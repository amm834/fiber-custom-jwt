package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-react-jwt/database"
	"go-react-jwt/models"
	"go-react-jwt/services"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func Register(c *fiber.Ctx) error {

	var request map[string]string

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid data"})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(request["password"]), 14)

	user := models.User{
		Name:     request["name"],
		Email:    request["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.Status(http.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {
	var request map[string]string

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	var user models.User

	database.DB.Where("email = ?", request["email"]).First(&user)

	if user.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(request["password"])); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	token, err := services.CreateToken(user.Id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error on login",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"token":   token,
		"expires": time.Now().Add(time.Hour * 24).Unix(), //24 hours
	})
}

func User(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims, err := services.Claims(tokenString)

	if err != nil {

		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})

	}

	var user models.User

	database.DB.Where("id = ?", claims["userId"]).First(&user)

	return c.Status(http.StatusOK).JSON(user)
	
}
