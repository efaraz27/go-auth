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

// Validate is a method that validates the user DTO
func (u *UserCreateDTO) Validate() error {
	return validateUserDTO(u)
}

// Validate is a method that validates the user DTO
func (u *UserUpdateDTO) Validate() error {
	return validateUserDTO(u)
}

func validateUserDTO(u interface{}) error {
	validate := validator.New()
	return validate.Struct(u)
}
