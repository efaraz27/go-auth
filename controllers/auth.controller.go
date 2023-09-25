package controllers

import (
	"github.com/efaraz27/go-auth/core"
	"github.com/efaraz27/go-auth/dtos"
	"github.com/efaraz27/go-auth/services"
	"github.com/go-playground/validator/v10"

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

	if err := ctx.BodyParser(registerDTO); err != nil {
		exception := core.NewBadRequestExceptionBuilder().WithMessage("Invalid JSON").Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	validator := validator.New()

	if err := validator.Struct(registerDTO); err != nil {
		exception := core.NewBadRequestExceptionBuilder().WithMessage("Invalid fields").Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	user, exception := c.service.Register(registerDTO.Email, registerDTO.Password, registerDTO.FirstName, registerDTO.LastName)

	if exception != nil {
		return ctx.Status(exception.Status).JSON(exception)
	}

	return ctx.Status(201).JSON(user)
}

// Login is a method that logs in a user
func (c *AuthController) Login(ctx *fiber.Ctx) error {

	loginDTO := new(dtos.LoginDTO)

	if err := ctx.BodyParser(loginDTO); err != nil {
		exception := core.NewBadRequestExceptionBuilder().WithMessage("Invalid JSON").Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	validator := validator.New()

	if err := validator.Struct(loginDTO); err != nil {
		exception := core.NewBadRequestExceptionBuilder().WithMessage("Invalid fields").Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	loginResponseDTO, exception := c.service.Login(loginDTO.Email, loginDTO.Password)

	if exception != nil {
		return ctx.Status(exception.Status).JSON(exception)
	}

	return ctx.Status(200).JSON(loginResponseDTO)
}
