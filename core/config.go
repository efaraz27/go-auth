// Description: This package is used to load the environment variables

package core

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config is a struct that defines the environment variables
type Config struct {
	AppName string
	AppPort int
	AppEnv  string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
}

// LoadConfig is a function that loads the environment variables
func LoadConfig() Config {
	// Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the environment variables
	appName := os.Getenv("APP_NAME")
	appPort, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	appEnv := os.Getenv("APP_ENV")

	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDBName := os.Getenv("POSTGRES_DBNAME")

	config := Config{
		AppName:          appName,
		AppPort:          appPort,
		AppEnv:           appEnv,
		PostgresHost:     postgresHost,
		PostgresPort:     postgresPort,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresDBName:   postgresDBName,
	}

	return config
}
