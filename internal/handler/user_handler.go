package handler

import (
	"github.com/Sudhanva05/Backend/internal/models"
	"github.com/Sudhanva05/Backend/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	validate *validator.Validate
	service  *service.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler() *UserHandler {
	return &UserHandler{
		validate: validator.New(),
		service:  service.NewUserService(),
	}
}

// GetUsers handles GET /users
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Get users handler working",
	})
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON body")
	}

	if err := h.validate.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	// Delegate business logic to service
	response := h.service.CreateUser(req)

	return c.Status(fiber.StatusCreated).JSON(response)
}
