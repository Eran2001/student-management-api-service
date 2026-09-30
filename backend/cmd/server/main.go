package main

import (
	"log"
	"net/http"
	"student-management-app/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	handlers.RegisterRoutes(mux)

	log.Println("Server starting on :8080...")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
