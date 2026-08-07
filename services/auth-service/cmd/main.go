package main

import (
	"log"
	"net/http"
	"os"
	"research-dev/services/auth-service/internal/config"
	"research-dev/services/auth-service/internal/handler"

	"github.com/Nerzal/gocloak/v13"
)

var mux = http.NewServeMux()
var cfg config.KeycloakConfig

func getEnv(key, defaultValue string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultValue
}
func init() {
	cfg = config.KeycloakConfig{
		BaseURL:      getEnv("KEYCLOAK_URL", "http://localhost:8080"),
		Realm:        getEnv("KEYCLOAK_REALM", "user-platform"),
		ClientID:     getEnv("KEYCLOAK_CLIENT_ID", "user-service"),
		ClientSecret: getEnv("KEYCLOAK_CLIENT_SECRET", "aogwy5ty6jPRLLPKbyfmqLaQyxWaYjKd"),
	}
}
func main() {

	gocloakClient := gocloak.NewClient(cfg.BaseURL)

	authHandler := handler.NewAuthHandler(gocloakClient, cfg)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth-service is running"))
	})

	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)

	log.Println("auth-service listening:8081")
	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}

}
