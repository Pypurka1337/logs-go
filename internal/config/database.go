package config

import (
	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST" default:"db"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"postgres"`
	DBName   string `envconfig:"DB_NAME" default:"postgres"`
	SSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`
}

func NewDatabaseConfig() (*DatabaseConfig, error) {
	var config DatabaseConfig
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *DatabaseConfig) ConnectionString() string {
	return "postgres://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.DBName + "?sslmode=" + c.SSLMode
}
