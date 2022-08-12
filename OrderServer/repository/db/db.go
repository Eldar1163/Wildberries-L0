package db

import (
	"OrderServer/logger"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var DB *gorm.DB

func DatabasebOpen() {
	err := godotenv.Load("../.env")
	if err != nil {
		logger.ErrorLogger.Fatalf("Cannot load environment variables. Err: %s", err)
	}
	connString := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("dbuser"),
		os.Getenv("pgsqlpass"),
		os.Getenv("dbname"))

	DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	if err != nil {
		logger.ErrorLogger.Panic("Cannot connect to database")
	}
}
