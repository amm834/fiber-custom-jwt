package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-react-jwt/controllers"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")

	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Get("/user", controllers.User)

}
