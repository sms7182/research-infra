package auth

import "github.com/Nerzal/gocloak/v13"

type KeycloakClient struct {
	Client       *gocloak.GoCloak
	Realm        string
	ClientID     string
	ClientSecret string
}

func NewKeycloak(
	url string,
	realm string,
	clientID string,
	clientSecret string,
) *KeycloakClient {
	return &KeycloakClient{
		Client:       gocloak.NewClient(url),
		Realm:        realm,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}
