package controllers

import (
	"golang-api/internal/models"
	"golang-api/internal/services"
	"golang-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// UserController handles user-related HTTP requests
type UserController struct {
	userService *services.UserService
}

// NewUserController creates a new UserController instance
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// GetAllUsers returns all users
func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	users := uc.userService.GetAllUsers()
	return utils.SuccessResponse(c, users, "Users retrieved successfully")
}

// GetUserByID returns a single user by ID
func (uc *UserController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, exists := uc.userService.GetUserByID(id)
	if !exists {
		return utils.NotFoundResponse(c, "User not found")
	}

	return utils.SuccessResponse(c, user, "User retrieved successfully")
}

// CreateUser creates a new user
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequestResponse(c, "Invalid request body")
	}

	// Validate input
	validationErrors := utils.ValidateCreateUser(req.Name, req.Email)
	if len(validationErrors) > 0 {
		return utils.BadRequestResponse(c, validationErrors[0])
	}

	// Create user
	user := uc.userService.CreateUser(req)

	return utils.CreatedResponse(c, user, "User created successfully")
}

// UpdateUser updates an existing user
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var req models.UpdateUserRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequestResponse(c, "Invalid request body")
	}

	// Update user
	user, exists := uc.userService.UpdateUser(id, req)
	if !exists {
		return utils.NotFoundResponse(c, "User not found")
	}

	return utils.SuccessResponse(c, user, "User updated successfully")
}

// DeleteUser deletes a user by ID
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	deleted := uc.userService.DeleteUser(id)
	if !deleted {
		return utils.NotFoundResponse(c, "User not found")
	}

	return utils.SuccessResponse(c, nil, "User deleted successfully")
}
