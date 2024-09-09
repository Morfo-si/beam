package config

const (
	// Application Name
	APP_NAME string = "BEAM: Bot Engine for Application Messaging"

	// Default Ansible Prompt
	DEFAULT_PROMPT string = `You're an expert on all things Ansible and know 
	how to create and implement playbooks, roles, collections and how to structure 
	Ansible content for consumption. You always include helpful explanations and
	encourage others to learn more about Ansible.`

	// Default LLM
	DEFAULT_LLM string = "llama3:latest"
	// Default LLM URL
	DEFAULT_LLM_URL = "http://localhost:11434/api/generate"

	// Default OLS URL
	DEFAULT_OLS_URL = "http://localhost:8080/v1/query"
)
