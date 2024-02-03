// Description: This package is used to load the environment variables

package core

import (
	"log"
	"os"
	"strconv"
	"time"

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

	RedisHost     string
	RedisPort     int
	RedisPassword string

	AmqpScheme   string
	AmqpHost     string
	AmqpPort     int
	AmqpUser     string
	AmqpPassword string
	AmqpQueue    string

	JwtSecret                   string
	AccessTokenExpDeltaSeconds  int
	RefreshTokenExpDeltaSeconds int

	EmailVerificationExpDeltaSeconds time.Duration
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

	redisHost := os.Getenv("REDIS_HOST")
	redisPort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	redisPassword := os.Getenv("REDIS_PASSWORD")

	amqpScheme := os.Getenv("AMQP_SCHEME")
	amqpHost := os.Getenv("AMQP_HOST")
	amqpPort, _ := strconv.Atoi(os.Getenv("AMQP_PORT"))
	amqpUser := os.Getenv("AMQP_USER")
	amqpPassword := os.Getenv("AMQP_PASSWORD")
	amqpQueue := os.Getenv("AMQP_QUEUE")

	jwtSecret := os.Getenv("JWT_SECRET")
	accessTokenExpDeltaSeconds, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXP_DELTA_SECONDS"))
	refreshTokenExpDeltaSeconds, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXP_DELTA_SECONDS"))

	EmailVerificationExpDeltaSeconds, _ := strconv.Atoi(os.Getenv("EMAIL_VERIFICATION_EXP_DELTA_SECONDS"))

	config := Config{
		AppName: appName,
		AppPort: appPort,
		AppEnv:  appEnv,

		PostgresHost:     postgresHost,
		PostgresPort:     postgresPort,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresDBName:   postgresDBName,

		RedisHost:     redisHost,
		RedisPort:     redisPort,
		RedisPassword: redisPassword,

		AmqpScheme:   amqpScheme,
		AmqpHost:     amqpHost,
		AmqpPort:     amqpPort,
		AmqpUser:     amqpUser,
		AmqpPassword: amqpPassword,
		AmqpQueue:    amqpQueue,

		JwtSecret:                   jwtSecret,
		AccessTokenExpDeltaSeconds:  accessTokenExpDeltaSeconds,
		RefreshTokenExpDeltaSeconds: refreshTokenExpDeltaSeconds,

		EmailVerificationExpDeltaSeconds: time.Duration(EmailVerificationExpDeltaSeconds) * time.Second,
	}

	return config
}
