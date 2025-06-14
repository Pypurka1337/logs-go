package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// HTTP server
	Port            string `envconfig:"PORT" default:"8080"`
	ReadTimeout     int    `envconfig:"READ_TIMEOUT" default:"15"`
	WriteTimeout    int    `envconfig:"WRITE_TIMEOUT" default:"15"`
	IdleTimeout     int    `envconfig:"IDLE_TIMEOUT" default:"60"`
	ShutdownTimeout int    `envconfig:"SHUTDOWN_TIMEOUT" default:"5"`

	// Application
	Environment    string `envconfig:"ENVIRONMENT" default:"development"`
	LogLevel       string `envconfig:"LOG_LEVEL" default:"info"`
	AllowedOrigins string `envconfig:"ALLOWED_ORIGINS" default:"*"`

	// CORS
	CORSAllowedMethods   string `envconfig:"CORS_ALLOWED_METHODS" default:"GET,POST,PUT,DELETE,OPTIONS"`
	CORSAllowedHeaders   string `envconfig:"CORS_ALLOWED_HEADERS" default:"Accept,Authorization,Content-Type,X-CSRF-Token"`
	CORSExposedHeaders   string `envconfig:"CORS_EXPOSED_HEADERS" default:"Link"`
	CORSAllowCredentials bool   `envconfig:"CORS_ALLOW_CREDENTIALS" default:"true"`
	CORSMaxAge           int    `envconfig:"CORS_MAX_AGE" default:"300"`
}

func Load() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
