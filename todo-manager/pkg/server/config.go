package server

import "os"

// Config holds server configuration
type Config struct {
	MysqlHost   string
	MysqlUser   string
	MysqlPass   string
	OcAgentHost string
}

// NewConfig loads config from environment variables
func NewConfig() *Config {
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost == "" {
		panic("Required environment variable 'MYSQL_HOST' not set")
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser == "" {
		panic("Required environment variable 'MYSQL_USER' not set")
	}
	mysqlPass := os.Getenv("MYSQL_PASS")
	if mysqlPass == "" {
		panic("Required environment variable 'MYSQL_PASS' not set")
	}
	ocAgentHost := os.Getenv("OC_AGENT_HOST")
	if ocAgentHost == "" {
		panic("Required environment variable 'OC_AGENT_HOST' not set")
	}

	return &Config{
		MysqlHost:   mysqlHost,
		MysqlUser:   mysqlUser,
		MysqlPass:   mysqlPass,
		OcAgentHost: ocAgentHost,
	}
}
