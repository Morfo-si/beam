package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ModelName string
	ModelAPI  string
	Host      string
	Port      string
	SSLMode   string
}

var configInstance *Config

// LoadConfig loads the environment variables and initializes the Config instance
func LoadConfig() *Config {
	if configInstance != nil {
		return configInstance
	}

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, falling back to system environment variables")
	}

	configInstance = &Config{
		Host:      getEnv("HOST", "localhost"),
		Port:      getEnv("PORT", "8081"),
		ModelName: getEnv("LLM_NAME", DEFAULT_LLM),
		ModelAPI:  getEnv("LLM_API", DEFAULT_LLM_URL),
		SSLMode:   getEnv("SSLMODE", "disable"),
	}

	return configInstance
}

// getEnv is a helper function to get an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
