package routes

import (
	"github.com/Sudhanva05/Backend/internal/handler"
	"github.com/gofiber/fiber/v2"
)

// RegisterUserRoutes registers all user-related routes
func RegisterUserRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	users := app.Group("/users")

	users.Get("/", userHandler.GetUsers)
	users.Get("/:id", userHandler.GetUserByID)
	users.Post("/", userHandler.CreateUser)
}
