package config

import "os"

type Config struct {
	Port          string
	DatabaseToken string
}

func NewConfig() *Config {
	// Consider using godotenv with .env file for this...
	p := os.Getenv("PORT")
	if p == "" {
		p = "8081"
	}

	dbToken := os.Getenv("FIRESTORE_ACCESS_TOKEN")
	if dbToken == "" {
		dbToken = "firestore_access_token.json"
	}

	return &Config{
		Port:          p,
		DatabaseToken: dbToken,
	}
}
