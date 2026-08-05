package handler

import (
	"research-dev/services/auth-service/internal/config"

	"github.com/Nerzal/gocloak/v13"
)

type AuthHandler struct {
	Gocloak *gocloak.GoCloak
	Config  config.KeycloakConfig
}

func NewAuthHandler(client *gocloak.GoCloak, cfg config.KeycloakConfig) *AuthHandler {
	return &AuthHandler{
		Gocloak: client,
		Config:  cfg,
	}
}
