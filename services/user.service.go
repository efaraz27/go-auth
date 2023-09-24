package services

import (
	"github.com/efaraz27/go-auth/core"
	"github.com/efaraz27/go-auth/models"
	"github.com/efaraz27/go-auth/repositories"
)

// UserService is a struct that defines the user service
type UserService struct {
	repository *repositories.UserRepository
}

// NewUserService is a function that returns a new user service
func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository}
}

// Create is a method that creates a new user
func (s *UserService) Create(user *models.User) (*models.User, *core.Exception, error) {
	// Check if user exists
	_, err := s.repository.FindByEmail(user.Email)

	if err == nil {
		return nil, core.NewBadRequestExceptionBuilder().WithMessage("User already exists").Build(), nil
	}

	return s.repository.Create(user)
}

// FindAll is a method that returns all users
func (s *UserService) FindAll() ([]models.User, error) {
	return s.repository.FindAll()
}

// FindByUUID is a method that returns a user by UUID
func (s *UserService) FindByUUID(uuid string) (*models.User, error) {
	return s.repository.FindByUUID(uuid)
}

// FindByEmail is a method that returns a user by email
func (s *UserService) FindByEmail(email string) (*models.User, error) {
	return s.repository.FindByEmail(email)
}

// Update is a method that updates a user
func (s *UserService) Update(user *models.User) (*models.User, error) {
	return s.repository.Update(user)
}

// Delete is a method that deletes a user
func (s *UserService) Delete(id int) error {
	return s.repository.Delete(id)
}
