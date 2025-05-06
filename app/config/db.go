package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, *sql.DB, error) {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_PORT")

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

	log.Println("Conntected to Database successfully....")

	return gDB, sqlDB, nil

}
