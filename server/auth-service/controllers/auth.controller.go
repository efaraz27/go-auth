package controllers

import (
	"github.com/efaraz27/go-auth/server/auth-service/core"
	"github.com/efaraz27/go-auth/server/auth-service/dtos"
	"github.com/efaraz27/go-auth/server/auth-service/services"

	"github.com/gofiber/fiber/v2"
)

// AuthController is a struct that defines the auth controller
type AuthController struct {
	service *services.AuthService
}

// NewAuthController is a function that returns a new auth controller
func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{service}
}

// Register is a method that registers a user
func (c *AuthController) Register(ctx *fiber.Ctx) error {

	registerDTO := new(dtos.RegisterDTO)

	err := core.ValidateBodyWithDTO(ctx, registerDTO)
	if err != nil {
		return err
	}

	user, exception := c.service.Register(ctx, registerDTO.Email, registerDTO.Password, registerDTO.FirstName, registerDTO.LastName)

	if exception != nil {
		return ctx.Status(exception.Status).JSON(exception)
	}

	return ctx.Status(201).JSON(user)
}

// Login is a method that logs in a user
func (c *AuthController) Login(ctx *fiber.Ctx) error {

	loginDTO := new(dtos.LoginDTO)

	err := core.ValidateBodyWithDTO(ctx, loginDTO)
	if err != nil {
		return err
	}

	loginResponseDTO, exception := c.service.Login(loginDTO.Email, loginDTO.Password)

	if exception != nil {
		return ctx.Status(exception.Status).JSON(exception)
	}

	return ctx.Status(200).JSON(loginResponseDTO)
}
