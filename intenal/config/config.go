// Package config contains config
package config

type Config struct {
	// ServerPort is the port for starting http-server
	ServerPort string `env:"SERVER_PORT"`
}