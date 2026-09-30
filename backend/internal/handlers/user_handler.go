package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"student-management-app/internal/models"
	"student-management-app/internal/repository"
)

// GetUser handles GET /user
func GetUser(write http.ResponseWriter, read *http.Request) {
	data := repository.GetUsers()

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusOK)
	err := json.NewEncoder(write).Encode(data)
	if err != nil {
		log.Println(err)
	}
}

// CreateUser handles POST /user
func CreateUser(write http.ResponseWriter, read *http.Request) {
	var input models.UserResponse

	err := json.NewDecoder(read.Body).Decode(&input)
	if err != nil {
		write.Header().Set("Content-Type", "application/json")
		write.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(write).Encode(map[string]string{"error": "Invalid JSON payload"})
		if err != nil {
			log.Println(err)
		}
		return
	}

	if input.Name == "" || input.Email == "" {
		write.Header().Set("Content-Type", "application/json")
		write.WriteHeader(http.StatusUnprocessableEntity)
		err := json.NewEncoder(write).Encode(map[string]string{"error": "Name and Email are required"})
		if err != nil {
			log.Println(err)
		}
		return
	}

	savedUser := repository.SaveUser(input)

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(write).Encode(savedUser)
	if err != nil {
		log.Println(err)
	}
}
