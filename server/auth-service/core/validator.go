package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateBodyWithDTO(ctx *fiber.Ctx, dto interface{}) error {
	if err := ctx.BodyParser(dto); err != nil {
		exception := NewBadRequestExceptionBuilder().WithMessage("Invalid JSON").Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	validator := validator.New()

	if err := validator.Struct(dto); err != nil {
		exception := NewBadRequestExceptionBuilder().WithMessage("Invalid fields").Build()
		return ctx.Status(exception.Status).JSON(exception)
	}

	return nil
}
