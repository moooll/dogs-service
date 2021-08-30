// Package config contains config
package config

// Config contains configuration strings for http server, database, etc.
type Config struct {
	// ServerPort is the port for starting http-server
	ServerPort string `env:"SERVER_PORT"`
}

// NewConfig returns new config
func NewConfig() *Config {
	return &Config{}
}