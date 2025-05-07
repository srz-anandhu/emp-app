package domain

import "time"

type Employee struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Name         string    `gorm:"column:name"`
	DOB          string    `gorm:"column:dob"`
	Email        string    `gorm:"column:email"`
	Password     string    `gorm:"column:password"`
	Phone        string    `gorm:"column:phone"`
	Address      string    `gorm:"column:address"`
	Salary       float64   `gorm:"column:salary"`
	Position     string    `gorm:"column:position"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
