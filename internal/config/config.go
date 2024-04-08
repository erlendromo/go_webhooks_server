package config

import "os"

type Config struct {
	Port string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	return &Config{
		Port: port,
	}
}
