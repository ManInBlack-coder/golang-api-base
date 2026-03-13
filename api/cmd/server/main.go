package main

import (
	"fmt"
	"log"

	"golang-api/internal/config"
	"golang-api/internal/controllers"
	"golang-api/internal/middleware"
	"golang-api/internal/routes"
	"golang-api/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Validate API key is set
	if cfg.APIKey == "" {
		log.Fatal("API_KEY environment variable is required")
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "Go Base API",
		ServerHeader: "Go-Base-API",
		ErrorHandler: customErrorHandler,
	})

	// Global middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, X-API-Key",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// API Key authentication middleware
	app.Use(middleware.APIKeyAuth(cfg.APIKey))

	// Initialize services
	userService := services.NewUserService()

	// Initialize controllers
	healthController := controllers.NewHealthController()
	userController := controllers.NewUserController(userService)

	// Setup routes
	routes.SetupRoutes(app, healthController, userController)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Route not found",
			"code":    404,
		})
	})

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("🚀 Server starting on http://localhost%s", addr)
	log.Printf("📋 Environment: %s", cfg.Environment)
	log.Printf("🔑 API Key authentication enabled")

	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// customErrorHandler handles Fiber errors
func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
		"code":    code,
	})
}
