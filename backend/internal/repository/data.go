package repository

import (
	"student-management-app/internal/models"
)

var users []models.UserResponse = []models.UserResponse{
	{ID: 1, Name: "John Doe", Email: "john_doe@example.com"},
}

func GetUsers() []models.UserResponse {
	return users
}

func SaveUser(newUser models.UserResponse) models.UserResponse {
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	return newUser
}
