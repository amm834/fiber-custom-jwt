package middewares

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// check authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	//token, err := jwt.ParseWithClaims(authHeader, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
	//	return []byte("secret"), nil
	//})

	return c.Next()
}
