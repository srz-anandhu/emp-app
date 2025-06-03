package config

import (
	"emp-app/app/domain"
	"emp-app/pkg/helpers/hash"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func AutoMigrateModels(db *gorm.DB) error {
	if err := db.AutoMigrate(&domain.Employee{}); err != nil {
		return fmt.Errorf("employee model migration failed due to : %v", err)
	}
	if err := db.AutoMigrate(&domain.Admin{}); err != nil {
		return fmt.Errorf("admin model migration failed due to : %v", err)
	}
	log.Println("migration successfull...")

	if err := CheckAndSeedAdmin(db); err != nil {
		return fmt.Errorf("admin creation failed due to : %v", err)
	}
	log.Println("Added admin successfully..!")
	return nil
}


// Seed admin in DB

func CheckAndSeedAdmin(db *gorm.DB) error {
	var count int64

	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := "adminpassword123"
		hashedPass, err := hash.HashPassword(password)
		if err != nil {
			return fmt.Errorf("check and admin creation error : %v", err)
		}
		admin := domain.Admin{
			ID: 1,
			Name: "admin",
			Email: "adminemail@gmail.com",
			Password: hashedPass,
		}
		db.Create(&admin)
	}
	return nil
}