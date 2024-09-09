package config

// Struct to define the input for the REST API
type ChatRequest struct {
	Prompt   string `json:"prompt"`
	Question string `json:"question"`
}

// Struct to define the input for OLS /v1/query API
type OLSRequest struct {
	Model          string `json:model`
	Provider       string `json:provider`
	Query          string `json:"query"`
	ConversationId string `json:conversation_id`
	// TODO Attachments
}
