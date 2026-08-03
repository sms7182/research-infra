package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Nerzal/gocloak/v13"
)

const (
	keycloakURL  = "http://localhost:8080"
	realm        = "master"
	clientID     = "user-service"
	clientSecret = "dQATtVAgYIicjH99sM9XLOmH279TAo2z" // Replace with your actual secret

	// Test user data
	testUsername = "testuser_go"
	testEmail    = "testuser_go@example.com"
	testPassword = "TestPassword123!"
)

func init() {

}

func main() {
	ctx := context.Background()
	client := gocloak.NewClient(keycloakURL)

	fmt.Println("🚀 Starting Keycloak Infrastructure Integration Test...")

	// -------------------------------------------------------------
	// STEP 1: Authenticate Backend Service Account (Get Admin Token)
	// -------------------------------------------------------------
	fmt.Print("1. Authenticating Go service account with Keycloak... ")
	adminToken, err := client.LoginClient(ctx, clientID, clientSecret, realm)
	if err != nil {
		log.Fatalf("\n❌ Client login failed: %v", err)
	}
	fmt.Println("SUCCESS!")

	// -------------------------------------------------------------
	// STEP 2: Create a Test User in Keycloak
	// -------------------------------------------------------------
	fmt.Printf("2. Creating test user '%s'... ", testUsername)
	newUser := gocloak.User{
		FirstName: gocloak.StringP("Test"),
		LastName:  gocloak.StringP("User"),
		Email:     gocloak.StringP(testEmail),
		Username:  gocloak.StringP(testUsername),
		Enabled:   gocloak.BoolP(true),
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				Type:      gocloak.StringP("password"),
				Value:     gocloak.StringP(testPassword),
				Temporary: gocloak.BoolP(false),
			},
		},
	}

	userID, err := client.CreateUser(ctx, adminToken.AccessToken, realm, newUser)
	if err != nil {
		fmt.Println("\n⚠️  User creation failed (It might already exist):", err)
	} else {
		fmt.Printf("SUCCESS! User ID: %s\n", userID)
	}

	// -------------------------------------------------------------
	// STEP 3: Authenticate as the newly created User
	// -------------------------------------------------------------
	fmt.Printf("3. Logging in as user '%s'... ", testUsername)
	userToken, err := client.Login(ctx, clientID, clientSecret, realm, testUsername, testPassword)
	if err != nil {
		log.Fatalf("\n❌ User authentication failed: %v", err)
	}
	fmt.Println("SUCCESS!")
	fmt.Printf("   -> Access Token (truncated): %s...\n", userToken.AccessToken[:30])

	// -------------------------------------------------------------
	// STEP 4: Validate / Introspect User's Access Token
	// -------------------------------------------------------------
	fmt.Print("4. Introspecting/Validating user token... ")
	rptResult, err := client.RetrospectToken(ctx, userToken.AccessToken, clientID, clientSecret, realm)
	if err != nil {
		log.Fatalf("\n❌ Token introspection failed: %v", err)
	}

	if rptResult.Active != nil && *rptResult.Active {
		fmt.Println("SUCCESS! Token is ACTIVE.")
	} else {
		log.Fatalf("\n❌ Token is INACTIVE or invalid.")
	}

	fmt.Println("\n🎉 ALL CHECKS PASSED! Your Keycloak infra is fully working with Go.")
}
