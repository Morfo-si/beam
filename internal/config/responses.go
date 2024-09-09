package config

// Request body for llama3:latest model API
type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

// Response from llama3:latest model
type OllamaResponse struct {
	Model    string `json:"model"`
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

// Reponse from OLS /v1/query API
type OLSResponse struct {
	ConversationId string `json:"conversation_id"`
	Response       string `json:"response"`
	Truncated      bool   `json:"truncated"`
	// TODO ReferencedDocuments
}
