package config

// Struct to define the input for the REST API
type ChatRequest struct {
	Prompt   string `json:"prompt"`
	Question string `json:"question"`
}
