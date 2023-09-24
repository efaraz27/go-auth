package routers

import (
	"github.com/efaraz27/go-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

// UserRouter is an interface that defines the user router
type UserRouter interface {
	SetupRoutes(app *fiber.App)
}

// userRouter is a struct that defines the user router
type userRouter struct {
	controller *controllers.UserController
}

// NewUserRouter is a function that returns a new user router
func NewUserRouter(controller *controllers.UserController) *userRouter {
	return &userRouter{controller}
}

// SetupRoutes is a method that sets up the routes
func (r *userRouter) SetupRoutes(app *fiber.App) {
	api := app.Group("/users")

	api.Get("/", r.controller.FindAll)
	api.Get("/uuid/:uuid", r.controller.FindByUUID)
	api.Get("/email/:email", r.controller.FindByEmail)
	api.Post("/", r.controller.Create)
	api.Put("/:id", r.controller.Update)
	api.Delete("/:id", r.controller.Delete)

}
