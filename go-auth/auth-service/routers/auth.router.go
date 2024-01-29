package routers

import (
	"github.com/efaraz27/go-auth/auth-service/controllers"

	"github.com/gofiber/fiber/v2"
)

// AuthRouter is an interface that defines the auth router
type AuthRouter interface {
	SetupRoutes(app *fiber.App)
}

// authRouter is a struct that defines the auth router
type authRouter struct {
	controller *controllers.AuthController
}

// NewAuthRouter is a function that returns a new auth router
func NewAuthRouter(controller *controllers.AuthController) *authRouter {
	return &authRouter{controller}
}

// SetupRoutes is a method that sets up the routes
func (r *authRouter) SetupRoutes(app *fiber.App) {
	api := app.Group("/auth")

	api.Post("/register", r.controller.Register)
	api.Post("/login", r.controller.Login)
}
