package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, *sql.DB, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("loading .env due to : %v", err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// postgres://postgres:password@localhost:5432/employee?sslmode=disable
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)

	gDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database : %v ", err)
	}

	if err := AutoMigrateModels(gDB); err != nil {
		log.Fatalf("automigration err : %v", err)
	}

	// Geting SQL DB object
	sqlDB, err := gDB.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Checking the connection is alive
	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Connected to Database successfully....")

	return gDB, sqlDB, nil

}
