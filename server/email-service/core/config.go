// Description: This package is used to load the environment variables

package core

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config is a struct that defines the environment variables
type Config struct {
	AmqpScheme   string
	AmqpHost     string
	AmqpPort     int
	AmqpUser     string
	AmqpPassword string
	AmqpQueue    string
}

// LoadConfig is a function that loads the environment variables
func LoadConfig() Config {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("failed to load the .env file")
	}

	// Get the environment variables
	amqpScheme := os.Getenv("AMQP_SCHEME")
	amqpHost := os.Getenv("AMQP_HOST")
	amqpPort, _ := strconv.Atoi(os.Getenv("AMQP_PORT"))
	amqpUser := os.Getenv("AMQP_USER")
	amqpPassword := os.Getenv("AMQP_PASSWORD")
	amqpQueue := os.Getenv("AMQP_QUEUE")

	return Config{
		AmqpScheme:   amqpScheme,
		AmqpHost:     amqpHost,
		AmqpPort:     amqpPort,
		AmqpUser:     amqpUser,
		AmqpPassword: amqpPassword,
		AmqpQueue:    amqpQueue,
	}
}
