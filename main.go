package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go-react-jwt/database"
	"go-react-jwt/routes"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	database.Connect()
}

func main() {

	app := fiber.New()
	app.Use(cors.New())

	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))
}
