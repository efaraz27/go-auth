package main

import (
	"log"

	"github.com/efaraz27/go-auth/server/auth-service/controllers"
	"github.com/efaraz27/go-auth/server/auth-service/core"
	"github.com/efaraz27/go-auth/server/auth-service/repositories"
	"github.com/efaraz27/go-auth/server/auth-service/repositories/store"
	"github.com/efaraz27/go-auth/server/auth-service/routers"
	"github.com/efaraz27/go-auth/server/auth-service/services"

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

	// Connect to redis
	redisClient := core.ConnetRedis(
		config.RedisHost,
		config.RedisPort,
		config.RedisPassword,
	)

	// Connet to RabbitMQ
	rabbitmq := core.NewRabbitMQ(
		config.AmqpScheme,
		config.AmqpHost,
		config.AmqpPort,
		config.AmqpUser,
		config.AmqpPassword,
	)

	// Create a queue
	queue, err := core.DeclareQueue(rabbitmq.Ch, config.AmqpQueue)
	if err != nil {
		panic("failed to declare a queue")
	}

	// setup repositories
	userRepository := repositories.NewUserRepository(db)

	// setup redis repositories
	emailVerificationRepository := store.NewEmailVerificationRepository(redisClient, config.EmailVerificationExpDeltaSeconds)

	// setup services
	emailService := services.NewEmailService(rabbitmq, queue.Name)
	tokenService := services.NewTokenService(config.JwtSecret, config.AccessTokenExpDeltaSeconds, config.RefreshTokenExpDeltaSeconds)
	userService := services.NewUserService(tokenService, userRepository, emailVerificationRepository)
	authService := services.NewAuthService(userService, tokenService, emailService)

	// setup controllers
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)

	// setup routes
	routers.NewUserRouter(userController).SetupRoutes(app)
	routers.NewAuthRouter(authController).SetupRoutes(app)

	app.Get("/api/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Go Auth Service",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
