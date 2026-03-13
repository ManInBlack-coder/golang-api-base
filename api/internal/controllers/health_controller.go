package controllers

import (
	"golang-api/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// HealthController handles health check requests
type HealthController struct{}

// NewHealthController creates a new HealthController instance
func NewHealthController() *HealthController {
	return &HealthController{}
}

// HealthCheck returns the API health status
func (h *HealthController) HealthCheck(c *fiber.Ctx) error {
	return utils.SuccessResponse(c, fiber.Map{
		"status":  "healthy",
		"version": "1.0.0",
	}, "API is running")
}
