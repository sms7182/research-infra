package handler

import (
	"context"
	"encoding/json"
	"net/http"
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

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	ctx := context.Background()
	jwt, err := h.Gocloak.Login(
		ctx,
		h.Config.ClientID,
		h.Config.ClientSecret,
		h.Config.Realm,
		req.Username,
		req.Password,
	)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	resp := TokenResponse{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
		TokenType:    jwt.TokenType,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
