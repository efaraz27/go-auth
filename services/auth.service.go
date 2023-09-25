package services

import (
	"math/rand"

	"github.com/efaraz27/go-auth/core"
	"github.com/efaraz27/go-auth/dtos"
	"github.com/efaraz27/go-auth/models"

	"golang.org/x/crypto/bcrypt"
)

// AuthService is a struct that defines the auth service
type AuthService struct {
	userService *UserService
	jwtService  *JWTService
}

// NewAuthService is a function that returns a new auth service
func NewAuthService(userService *UserService, jwtService *JWTService) *AuthService {
	return &AuthService{userService, jwtService}
}

// Register is a method that registers a user
func (s *AuthService) Register(email string, password string, firstName string, lastName string) (*models.User, *core.Exception) {

	hashedPassword, err := hashPassword(password)

	if err != nil {
		exception := core.NewInternalServerErrorExceptionBuilder().WithMessage("Could not register user").Build()
		return nil, exception
	}

	user, exception := s.userService.Create(email, hashedPassword, firstName, lastName)

	if exception != nil {
		return nil, exception
	}

	return user, nil

}

// Login is a method that logs in a user
func (s *AuthService) Login(email string, password string) (*dtos.LoginResponseDTO, *core.Exception) {

	user, exception := s.userService.FindByEmail(email)

	if exception != nil {
		return nil, exception
	}

	if !comparePassword(user.Password, password) {
		exception := core.NewUnauthorizedExceptionBuilder().WithMessage("Invalid email or password").Build()
		return nil, exception
	}

	accessToken, exception := s.jwtService.GenerateAccessToken(user.Uuid)

	if exception != nil {
		return nil, exception
	}

	refreshToken, exception := s.jwtService.GenerateRefreshToken(user.Uuid)

	if exception != nil {
		return nil, exception
	}

	loginResponseDTO := &dtos.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return loginResponseDTO, nil

}

func hashPassword(password string) (string, error) {

	// random cost between bcrypt.MinCost and bcrypt.MaxCost
	cost := rand.Intn(bcrypt.MaxCost-bcrypt.MinCost) + bcrypt.MinCost

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", err
	}

	return string(bytes), nil

}

func comparePassword(hashedPassword string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil

}
