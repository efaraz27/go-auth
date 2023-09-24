package dtos

import "github.com/go-playground/validator/v10"

// UserCreateDTO is a struct that defines the user create DTO
type UserCreateDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

// UserUpdateDTO is a struct that defines the user update DTO
type UserUpdateDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

// Validate is a method that validates the user create DTO
func (u *UserCreateDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// Validate is a method that validates the user update DTO
func (u *UserUpdateDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
