package core

import (
	"fmt"
	"log"

	"github.com/efaraz27/go-auth/auth-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB is a function that connects to the database
func ConnectDB(
	host string,
	port int,
	user string,
	password string,
	dbname string,
) *gorm.DB {

	// Connect to the default 'postgres' database to check existence of our target database
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=postgres",
		host,
		port,
		user,
		password,
	)
	defaultDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the default postgres database")
	}

	// Check if the desired database exists
	var count int64
	defaultDB.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", dbname).Scan(&count)
	if count == 0 {
		log.Printf("Database %s does not exist. Creating...", dbname)
		defaultDB.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbname))
	}

	// Close the connection to default 'postgres' database
	sqlDB, err := defaultDB.DB()
	if err != nil {
		panic("failed to get database from GORM connection")
	}
	sqlDB.Close()

	// Connect to our target database
	dsn = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		host,
		port,
		user,
		password,
		dbname,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		panic(err)
	}

	return db
}
