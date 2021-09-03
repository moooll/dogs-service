// Package config contains config
package config

// Config contains configuration strings for http server, database, etc.
type Config struct {
	// ServerPort is the port for starting http-server
	ServerPort string `env:"SERVER_PORT"`

	// DatabaseURI is the database connection string
	DatabaseURI string `env:"DATABASE_URI"`

	// JWTSecret is the secret key for creating JWT
	JWTSecret string `env:"JWT_SECRET"`
}

// NewConfig returns new config
func NewConfig() *Config {
	return &Config{}
}