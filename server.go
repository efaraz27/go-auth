package main

import (
	"log"

	"github.com/efaraz27/go-auth/controllers"
	"github.com/efaraz27/go-auth/core"
	"github.com/efaraz27/go-auth/repositories"
	"github.com/efaraz27/go-auth/routers"
	"github.com/efaraz27/go-auth/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	// Load the environment variables
	config := core.LoadConfig()

	app := fiber.New()
	app.Use(logger.New())

	// Connect to the database
	db := core.ConnectDB(
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDBName,
	)

	// setup repositories
	userRepository := repositories.NewUserRepository(db)

	// setup services
	userService := services.NewUserService(userRepository)
	jwtService := services.NewJWTService(config.JwtSecret, config.AccessTokenExpDeltaSeconds, config.RefreshTokenExpDeltaSeconds)
	authService := services.NewAuthService(userService, jwtService)

	// setup controllers
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)

	// setup routes
	routers.NewUserRouter(userController).SetupRoutes(app)
	routers.NewAuthRouter(authController).SetupRoutes(app)

	app.Get("/api/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
