package services

import (
	"crypto/rand"
	"time"

	"github.com/efaraz27/go-auth/core"
	"github.com/google/uuid"

	"github.com/golang-jwt/jwt"
)

const BASE_62_CHARS = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// TokenService is a struct that defines the token service
type TokenService struct {
	jwtSecret                   string
	accessTokenExpDeltaSeconds  int
	refreshTokenExpDeltaSeconds int
}

// NewTokenService is a function that returns a new token service
func NewTokenService(jwtSecret string, accessTokenExpDeltaSeconds int, refreshTokenExpDeltaSeconds int) *TokenService {
	return &TokenService{
		jwtSecret:                   jwtSecret,
		accessTokenExpDeltaSeconds:  accessTokenExpDeltaSeconds,
		refreshTokenExpDeltaSeconds: refreshTokenExpDeltaSeconds,
	}
}

// GenerateAccessToken is a method that generates an access token
func (j *TokenService) GenerateAccessToken(uuid uuid.UUID) (string, *core.Exception) {

	claims := jwt.MapClaims{
		"uuid": uuid,
		"exp":  time.Now().Add(time.Second * time.Duration(j.accessTokenExpDeltaSeconds)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(j.jwtSecret))

	if err != nil {
		exception := core.NewInternalServerErrorExceptionBuilder().WithMessage("Could not generate access token").Build()
		return "", exception
	}

	return accessToken, nil

}

// GenerateRefreshToken is a method that generates a refresh token
func (j *TokenService) GenerateRefreshToken(uuid uuid.UUID) (string, *core.Exception) {
	claims := jwt.MapClaims{
		"uuid": uuid,
		"exp":  time.Now().Add(time.Second * time.Duration(j.refreshTokenExpDeltaSeconds)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := token.SignedString([]byte(j.jwtSecret))

	if err != nil {
		exception := core.NewInternalServerErrorExceptionBuilder().WithMessage("Could not generate refresh token").Build()
		return "", exception
	}

	return refreshToken, nil
}

// GenerateRandomToken is a method that generates a random token
func (j *TokenService) GenerateRandomToken() (string, *core.Exception) {
	b := make([]byte, 32)
	_, err := rand.Read(b)

	if err != nil {
		exception := core.NewInternalServerErrorExceptionBuilder().WithMessage("Could not generate random token").Build()
		return "", exception
	}

	token := ""

	for _, byteVal := range b {
		token += string(BASE_62_CHARS[byteVal%62])
	}

	return token, nil

}
