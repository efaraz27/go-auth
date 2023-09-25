package services

import (
	"time"

	"github.com/efaraz27/go-auth/core"
	"github.com/google/uuid"

	"github.com/golang-jwt/jwt"
)

// JWT is a struct that defines the JWT
type JWTService struct {
	jwtSecret                   string
	accessTokenExpDeltaSeconds  int
	refreshTokenExpDeltaSeconds int
}

// NewJWT is a function that returns a new JWT
func NewJWTService(jwtSecret string, accessTokenExpDeltaSeconds int, refreshTokenExpDeltaSeconds int) *JWTService {
	return &JWTService{
		jwtSecret:                   jwtSecret,
		accessTokenExpDeltaSeconds:  accessTokenExpDeltaSeconds,
		refreshTokenExpDeltaSeconds: refreshTokenExpDeltaSeconds,
	}
}

// GenerateAccessToken is a method that generates an access token
func (j *JWTService) GenerateAccessToken(uuid uuid.UUID) (string, *core.Exception) {

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
func (j *JWTService) GenerateRefreshToken(uuid uuid.UUID) (string, *core.Exception) {
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
