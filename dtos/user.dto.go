package dtos

import (
	"github.com/efaraz27/go-auth/models"
)

// UserUpdateDTO is a struct that defines the user update DTO
type UserUpdateDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

type UserResponseDTO = models.User
type UserListResponseDTO = []models.User
