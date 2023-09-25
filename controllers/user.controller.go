package controllers

import (
	"github.com/efaraz27/go-auth/core"
	"github.com/efaraz27/go-auth/dtos"
	"github.com/efaraz27/go-auth/models"
	"github.com/efaraz27/go-auth/services"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

// UserController is a struct that defines the user controller
type UserController struct {
	service *services.UserService
}

// NewUserController is a function that returns a new user controller
func NewUserController(service *services.UserService) *UserController {
	return &UserController{service}
}

// FindAll is a method that returns all users
func (c *UserController) FindAll(ctx *fiber.Ctx) error {

	users, exception := c.service.FindAll()

	if exception != nil {
		return ctx.Status(exception.Status).JSON(exception)
	}

	return ctx.Status(200).JSON(users)
}

// FindByUUID is a method that returns a user by UUID
func (c *UserController) FindByUUID(ctx *fiber.Ctx) error {

	uuid, err := uuid.Parse(ctx.Params("uuid"))

	if err != nil {
		exception := core.NewBadRequestExceptionBuilder().WithMessage("Invalid UUID").Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	user, exception := c.service.FindByUUID(uuid)

	if exception != nil {
		return ctx.Status(exception.Status).JSON(exception)
	}

	return ctx.Status(200).JSON(user)
}

// FindByEmail is a method that returns a user by email
func (c *UserController) FindByEmail(ctx *fiber.Ctx) error {

	user, exception := c.service.FindByEmail(ctx.Params("email"))

	if exception != nil {
		return ctx.Status(exception.Status).JSON(exception)
	}

	return ctx.Status(200).JSON(user)
}

// Create is a method that creates a new user
func (c *UserController) Create(ctx *fiber.Ctx) error {
	user := new(dtos.UserCreateDTO)

	if err := ctx.BodyParser(user); err != nil {
		exception := core.NewBadRequestExceptionBuilder().WithMessage("Unable to parse JSON").Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	if err := user.Validate(); err != nil {
		exception := core.NewBadRequestExceptionBuilder().WithMessage("Invalid request body").WithPayload(err.Error()).Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	createdUser, exception := c.service.Create(&newUser)

	if exception != nil {
		return ctx.Status(exception.Status).JSON(exception)
	}

	return ctx.Status(201).JSON(createdUser)
}

// Update is a method that updates a user
func (c *UserController) Update(ctx *fiber.Ctx) error {
	user := new(dtos.UserUpdateDTO)

	if err := ctx.BodyParser(user); err != nil {
		exception := core.NewBadRequestExceptionBuilder().WithMessage("Unable to parse JSON").Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	if err := user.Validate(); err != nil {
		exception := core.NewBadRequestExceptionBuilder().WithMessage("Invalid request body").WithPayload(err.Error()).Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	userUpdate := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	updatedUser, exception := c.service.Update(&userUpdate)

	if exception != nil {
		return ctx.Status(exception.Status).JSON(exception)
	}

	return ctx.Status(200).JSON(updatedUser)
}

// Delete is a method that deletes a user
func (c *UserController) Delete(ctx *fiber.Ctx) error {

	uuid, err := uuid.Parse(ctx.Params("uuid"))

	if err != nil {
		exception := core.NewBadRequestExceptionBuilder().WithMessage("Invalid UUID").Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	if exception := c.service.Delete(uuid); exception != nil {
		return ctx.Status(exception.Status).JSON(exception)
	}

	return ctx.Status(200).JSON(nil)
}
