package main

import (
	"log"

	"github.com/Sudhanva05/Backend/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Backend server running")
	})

	// Register user routes
	routes.RegisterUserRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
