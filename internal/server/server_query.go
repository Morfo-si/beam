package server

import (
	"fmt"

	"github.com/morfo-si/beam/internal/config"

	"github.com/gofiber/fiber/v3"
)

// ACEServer struct
func (ace *ACEServer) Query(c fiber.Ctx) error {
	// Parse the request body
	var chatRequest config.ChatRequest
	if err := c.Bind().Body(&chatRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate inputs
	if chatRequest.Prompt == "" {
		chatRequest.Prompt = config.DEFAULT_PROMPT
	}
	if chatRequest.Question == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Prompt and question are required"})
	}

	// Combine prompt and question
	fullPrompt := fmt.Sprintf("%s\nQuestion: %s", chatRequest.Prompt, chatRequest.Question)

	// Send the prompt to the LLM
	concatenatedResponse, err := queryLLM(fullPrompt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return the concatenated response as JSON
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": concatenatedResponse})
}
