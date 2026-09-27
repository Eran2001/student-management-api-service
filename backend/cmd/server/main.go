package main

import (
	"fmt"
	"net/http"
)

// main function to initialize and start a basic go server
func main() {
	fmt.Println("Starting server on :8080")

	// need proper JSON response
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"message": "Hello, World!"}`)
	})

	// need proper POST API endpoint
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"message": "POST request received"}`)
	})

	http.ListenAndServe(":8080", nil)
}
