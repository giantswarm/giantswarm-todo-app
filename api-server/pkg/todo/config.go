package todo

import "os"

// Config holds server configuration
type Config struct {
	TodoURL     string
	OcAgentHost string
}

// NewConfig loads config from environment variables
func NewConfig() *Config {
	todoURL := os.Getenv("TODO_URL")
	if todoURL == "" {
		panic("Required environment variable 'TODO_URL' not set")
	}
	ocAgentHost := os.Getenv("OC_AGENT_HOST")
	if ocAgentHost == "" {
		panic("Required environment variable 'OC_AGENT_HOST' not set")
	}

	return &Config{
		TodoURL:     todoURL,
		OcAgentHost: ocAgentHost,
	}
}
