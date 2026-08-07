package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Nerzal/gocloak/v13"
)

type RegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	adminToken, err := h.Gocloak.LoginClient(
		ctx,
		h.Config.ClientID,
		h.Config.ClientSecret,
		h.Config.Realm,
	)
	if err != nil {
		http.Error(w, "Failed to authenticate with Keycloak Admin", http.StatusInternalServerError)
		return
	}

	enabled := true

	keyCloakUser := gocloak.User{
		Username:        &req.Username,
		Email:           &req.Email,
		FirstName:       &req.FirstName,
		LastName:        &req.LastName,
		Enabled:         &enabled,
		EmailVerified:   &enabled,
		RequiredActions: &[]string{},
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				Type:      gocloak.StringP("password"),
				Value:     &req.Password,
				Temporary: gocloak.BoolP(false),
			},
		},
	}

	userID, err := h.Gocloak.CreateUser(ctx, adminToken.AccessToken, h.Config.Realm, keyCloakUser)
	if err != nil {
		http.Error(w, "keycloak user creaiton failed:"+err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
		"user_id": userID,
	})

}
