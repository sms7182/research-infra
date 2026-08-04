package main

import (
	"log"
	"net/http"
)

var mux = http.NewServeMux()

func main() {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth-service is running"))
	})
	log.Println("auth-service listening:8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}
