package controllers

import (
	"strconv"

	"github.com/efaraz27/go-auth/dtos"
	"github.com/efaraz27/go-auth/models"
	"github.com/efaraz27/go-auth/services"

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
	users, err := c.service.FindAll()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Unable to fetch users",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   users,
	})
}

// FindByUUID is a method that returns a user by UUID
func (c *UserController) FindByUUID(ctx *fiber.Ctx) error {
	user, err := c.service.FindByUUID(ctx.Params("uuid"))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Unable to fetch user",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}

// FindByEmail is a method that returns a user by email
func (c *UserController) FindByEmail(ctx *fiber.Ctx) error {
	user, err := c.service.FindByEmail(ctx.Params("email"))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Unable to fetch user",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}

// Create is a method that creates a new user
func (c *UserController) Create(ctx *fiber.Ctx) error {
	user := new(dtos.UserCreateDTO)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Unable to parse JSON",
		})
	}

	err := user.Validate()

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	createdUser, exception, err := c.service.Create(&newUser)

	if exception != nil {
		return ctx.Status(exception.Status).JSON(exception)
	}

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Unable to create user",
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"status": "success",
		"data":   createdUser,
	})
}

// Update is a method that updates a user
func (c *UserController) Update(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Unable to parse JSON",
		})
	}

	updatedUser, err := c.service.Update(user)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Unable to update user",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   updatedUser,
	})
}

// Delete is a method that deletes a user
func (c *UserController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Unable to parse ID",
		})
	}

	if err := c.service.Delete(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Unable to delete user",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully deleted",
	})
}
