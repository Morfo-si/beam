package config

import (
	"errors"
	"log"

	"github.com/joho/godotenv"
)

// Define a custom type for the engine
type EngineType string

// Define constants for the allowed engine values
const (
	EngineBeam EngineType = "beam"
	EngineOLS  EngineType = "ols"
)

// Config struct with the new engine field
type Config struct {
	ModelName string
	ModelAPI  string
	Host      string
	Port      string
	Engine    EngineType
}

// Function to validate the engine field
func (c *Config) Validate() error {
	if c.Engine != EngineBeam && c.Engine != EngineOLS {
		return errors.New("invalid engine type: must be 'beam' or 'ols'")
	}
	return nil
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
		Engine:    getEngineEnv("ENGINE", EngineBeam),
	}
	return configInstance
}
