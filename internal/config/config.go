package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port          string
	DatabaseToken string
}

func NewConfig() (*Config, error) {
	p := os.Getenv("PORT")
	if p == "" {
		return nil, fmt.Errorf("unable to set server-port")
	}

	dbToken := os.Getenv("FIRESTORE_ACCESS_TOKEN")
	if dbToken == "" {
		return nil, fmt.Errorf("unable to load firestore access token")
	}

	return &Config{
		Port:          p,
		DatabaseToken: dbToken,
	}, nil
}
