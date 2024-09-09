package server

import (
	"github.com/morfo-si/beam/internal/config"

	"github.com/gofiber/fiber/v3"
)

// ACEServer struct
func (ace *ACEServer) QueryOLS(c fiber.Ctx) error {
	// Parse the request body
	var chatRequest config.ChatRequest
	if err := c.Bind().Body(&chatRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Note: Prompt is ignored when OLS is used.

	if chatRequest.Question == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Question is required for OLS"})
	}

	// Send the question to the OLS
	concatenatedResponse, err := queryOLS(chatRequest.Question)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return the concatenated response as JSON
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": concatenatedResponse})
}
