package routes

import (
	"golang-api/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all API routes
func SetupRoutes(
	app *fiber.App,
	healthController *controllers.HealthController,
	userController *controllers.UserController,
) {
	// API group
	api := app.Group("/api")

	// Health check (no authentication required)
	api.Get("/health", healthController.HealthCheck)

	// User routes (authentication required via middleware)
	users := api.Group("/users")
	users.Get("/", userController.GetAllUsers)
	users.Get("/:id", userController.GetUserByID)
	users.Post("/", userController.CreateUser)
	users.Put("/:id", userController.UpdateUser)
	users.Delete("/:id", userController.DeleteUser)
}
