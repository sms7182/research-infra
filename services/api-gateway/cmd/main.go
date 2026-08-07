package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var authServiceURL string
var targetURL string
var proxy *httputil.ReverseProxy
var mux *http.ServeMux

func init() {
	authServiceURL = os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		authServiceURL = "http://localhost:8082"
	}

}
func main() {
	targetURL, err := url.Parse(authServiceURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse auth service URL: %v", err))
	}
	proxy = httputil.NewSingleHostReverseProxy(targetURL)

	authProxy := httputil.NewSingleHostReverseProxy(targetURL)

	authHandler := http.StripPrefix("/api/v1/auth", authProxy)
	mux = http.NewServeMux()
	mux.Handle("/api/v1/auth/", authHandler)
	fmt.Println("API gateway running in on port :8000...")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		panic(err)
	}
}
