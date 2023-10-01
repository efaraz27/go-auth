package services

import (
	"context"

	"github.com/efaraz27/go-auth/core"
	"github.com/efaraz27/go-auth/models"
	"github.com/efaraz27/go-auth/repositories"
	"github.com/efaraz27/go-auth/repositories/store"
	"github.com/google/uuid"
)

// UserService is a struct that defines the user service
type UserService struct {
	tokenService                *TokenService
	repository                  *repositories.UserRepository
	emailVerificationRepository *store.EmailVerificationRepository
}

// NewUserService is a function that returns a new user service
func NewUserService(tokenService *TokenService, repository *repositories.UserRepository, emailVerificationRepository *store.EmailVerificationRepository) *UserService {
	return &UserService{tokenService, repository, emailVerificationRepository}
}

// FindAll is a method that returns all users
func (s *UserService) FindAll() ([]models.User, *core.Exception) {

	users, err := s.repository.FindAll()

	if err != nil {
		return nil, core.NewInternalServerErrorExceptionBuilder().WithMessage("Unable to fetch users").Build()
	}

	return users, nil
}

// FindByUUID is a method that returns a user by UUID
func (s *UserService) FindByUUID(uuid uuid.UUID) (*models.User, *core.Exception) {

	user, err := s.repository.FindByUUID(uuid)

	if err != nil {
		return nil, core.NewInternalServerErrorExceptionBuilder().WithMessage("Unable to fetch user").Build()
	}

	return user, nil
}

// FindByEmail is a method that returns a user by email
func (s *UserService) FindByEmail(email string) (*models.User, *core.Exception) {

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return nil, core.NewInternalServerErrorExceptionBuilder().WithMessage("Unable to fetch user").Build()
	}

	return user, nil
}

// Create is a method that creates a new user
func (s *UserService) Create(email string, password string, firstName string, lastName string) (*models.User, *core.Exception) {

	// Check if user exists
	_, err := s.repository.FindByEmail(email)

	if err == nil {
		return nil, core.NewBadRequestExceptionBuilder().WithMessage("User already exists").Build()
	}

	user := &models.User{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

	user, err = s.repository.Create(user)

	if err != nil {
		return nil, core.NewInternalServerErrorExceptionBuilder().WithMessage("Unable to create user").Build()
	}

	return user, nil

}

// Update is a method that updates a user
func (s *UserService) Update(email string, firstName string, lastName string) (*models.User, *core.Exception) {
	// Check if user exists
	_, err := s.repository.FindByEmail(email)

	if err != nil {
		return nil, core.NewBadRequestExceptionBuilder().WithMessage("User does not exist").Build()
	}

	user := &models.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	user, err = s.repository.Update(user)

	if err != nil {
		return nil, core.NewInternalServerErrorExceptionBuilder().WithMessage("Unable to update user").Build()
	}

	return user, nil
}

// Delete is a method that deletes a user
func (s *UserService) Delete(uuid uuid.UUID) *core.Exception {
	// Check if user exists
	_, err := s.repository.FindByUUID(uuid)

	if err != nil {
		return core.NewBadRequestExceptionBuilder().WithMessage("User does not exist").Build()
	}

	err = s.repository.Delete(uuid)

	if err != nil {
		return core.NewInternalServerErrorExceptionBuilder().WithMessage("Unable to delete user").Build()
	}

	return nil
}

// VerifyEmail is a method that sends an email verification email
func (s *UserService) VerifyEmail(ctx context.Context, email string) *core.Exception {
	// Check if user exists
	_, err := s.repository.FindByEmail(email)

	if err != nil {
		return core.NewBadRequestExceptionBuilder().WithMessage("User does not exist").Build()
	}

	// Generate email verification token
	token, exception := s.tokenService.GenerateRandomToken()

	if exception != nil {
		return exception
	}

	// Save email verification token

	err = s.emailVerificationRepository.Create(ctx, email, token)

	if err != nil {
		return core.NewInternalServerErrorExceptionBuilder().WithMessage("Unable to generate email verification token").Build()
	}

	// Send email verification email
	// TODO: Implement email verification email

	return nil
}
