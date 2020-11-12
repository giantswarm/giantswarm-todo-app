package todo

import (
	"os"
	"strconv"
)

// Config holds server configuration
type Config struct {
	TodoURL        string
	OcAgentHost    string
	EnableFailures bool
	EnableTracing  bool
}

// NewConfig loads config from environment variables
func NewConfig() *Config {
	todoURL := os.Getenv("TODO_URL")
	if todoURL == "" {
		panic("Required environment variable 'TODO_URL' not set")
	}
	ocAgentHost := os.Getenv("OC_AGENT_HOST")
	boolEnableFailures := false
	enableFailures := os.Getenv("ENABLE_FAILURES")
	if enableFailures != "" {
		if b, err := strconv.ParseBool(enableFailures); err == nil && b {
			boolEnableFailures = true
		}
	}
	boolEnableTracing := false
	enableTracing := os.Getenv("ENABLE_TRACING")
	if enableTracing != "" {
		if b, err := strconv.ParseBool(enableTracing); err == nil {
			boolEnableTracing = b
		}
	}
	if boolEnableTracing && ocAgentHost == "" {
		panic("Required environment variable 'OC_AGENT_HOST' not set")
	}

	return &Config{
		TodoURL:        todoURL,
		OcAgentHost:    ocAgentHost,
		EnableFailures: boolEnableFailures,
		EnableTracing:  boolEnableTracing,
	}
}
