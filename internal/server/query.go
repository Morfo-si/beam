package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/morfo-si/beam/internal/config"
)

// Function to query the LLM API
func queryLLM(fullPrompt string) (string, error) {
	// Define the request body for the LLM API
	requestBody := config.OllamaRequest{
		Model:  config.LoadConfig().ModelName,
		Prompt: fullPrompt,
	}

	// Marshal the request body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshalling request: %v", err)
	}

	// Send the request to the local LLM API
	resp, err := http.Post(config.LoadConfig().ModelAPI, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error sending request to LLM: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading LLM response: %v", err)
	}

	// Parse and concatenate the streaming response from LLaMA
	var responseParts []string
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)

	// Iterate through the streaming response and concatenate
	for decoder.More() {
		var part config.OllamaResponse
		if err := decoder.Decode(&part); err != nil {
			return "", fmt.Errorf("error decoding LLM response: %v", err)
		}
		responseParts = append(responseParts, part.Response)

		if part.Done {
			break
		}
	}

	// Return the concatenated response
	return concatResponses(responseParts), nil
}

// Helper function to concatenate responses
func concatResponses(parts []string) string {
	var finalResponse string
	for _, part := range parts {
		finalResponse += part
	}
	return finalResponse
}
