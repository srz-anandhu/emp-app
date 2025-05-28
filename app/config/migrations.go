package config

import (
	"emp-app/app/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func AutoMigrateModels(db *gorm.DB) error {
	if err := db.AutoMigrate(&domain.Employee{}); err != nil {
		return fmt.Errorf("employee model migration failed due to : %v", err)
	}
	log.Println("migration successfull...")
	return nil
}


// Seed admin in DB