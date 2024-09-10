package config

import (
	"fmt"
	"os"
)

// getEnv is a helper function to get an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Helper function to get Engine from environment with validation
func getEngineEnv(key string, defaultValue EngineType) EngineType {
	value := os.Getenv(key)

	switch value {
	case string(EngineBeam), string(EngineOLS):
		return EngineType(value)
	default:
		fmt.Printf("Invalid or missing engine value. Defaulting to %s\n", defaultValue)
		return defaultValue
	}
}
