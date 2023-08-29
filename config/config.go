package config

import (
	"fmt"

	"github.com/gofor-little/env"
)

type Config struct {
	DBConfig     DBConfig
	ServerConfig ServerConfig
}

type DBConfig struct {
	DB_NAME     string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
}

type ServerConfig struct {
	SERVER_HOST      string
	API_V1_BASE_PATH string
	API_V2_BASE_PATH string
}

func Load() (*Config, error) {
	// Load an .env file and set the key-value pairs as environment variables.
	if err := env.Load(".env"); err != nil {
		return nil, err
	}

	// load server config variable values
	serverIp := env.Get("SERVER_IP", "localhost")
	serverPort := env.Get("SERVER_PORT", "8080")

	c := Config{
		ServerConfig: ServerConfig{
			SERVER_HOST:      fmt.Sprintf("%s:%s", serverIp, serverPort),
			API_V1_BASE_PATH: "/api/v1",
			API_V2_BASE_PATH: "/api/v2",
		},
		DBConfig: DBConfig{
			DB_HOST:     env.Get("DB_HOST", "localhost"),
			DB_NAME:     env.Get("DB_NAME", "postgres"),
			DB_USER:     env.Get("DB_USER", "admin"),
			DB_PASSWORD: env.Get("DB_PASSWORD", "admin"),
		},
	}

	return &c, nil
}
