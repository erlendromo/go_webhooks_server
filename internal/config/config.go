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
	port, present := os.LookupEnv("PORT")
	if !present {
		port = "8081"
	}

	dbToken, present := os.LookupEnv("FIRESTORE_ACCESS_TOKEN")
	if !present {
		return nil, fmt.Errorf("unable to load firestore access token")
	}

	return &Config{
		Port:          port,
		DatabaseToken: dbToken,
	}, nil
}
