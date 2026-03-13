package middleware

import (
	"golang-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// APIKeyAuth creates a middleware that validates API key from X-API-Key header
func APIKeyAuth(validAPIKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Skip authentication for health check endpoint
		if c.Path() == "/api/health" {
			return c.Next()
		}

		// Get API key from header
		apiKey := c.Get("X-API-Key")

		// Check if API key is provided
		if apiKey == "" {
			return utils.UnauthorizedResponse(c, "API key is required. Please provide X-API-Key header")
		}

		// Validate API key
		if apiKey != validAPIKey {
			return utils.UnauthorizedResponse(c, "Invalid API key")
		}

		// API key is valid, continue to next handler
		return c.Next()
	}
}
