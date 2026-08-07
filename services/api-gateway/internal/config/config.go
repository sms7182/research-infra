package config

import "os"

type Config struct {
	Port           string
	AuthServiceURL string
}

func LoadConfig() Config {
	authURL := os.Getenv("AUTH_SERVICE_URL")
	if authURL == "" {
		authURL = "http://localhost:8082"
	}
	port := os.Getenv("GATEWAY_PORT")
	if port == "" {
		port = "8000"
	}
	return Config{
		Port:           port,
		AuthServiceURL: authURL,
	}

}
