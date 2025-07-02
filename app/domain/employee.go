package domain

import (
	"encoding/json"
	"time"
)

type Employee struct {
	ID         int         `gorm:"primaryKey;autoIncrement"`
	EmployeeID string      `gorm:"column:employee_id;unique"`
	FullName   string      `gorm:"column:name"`
	DOB        string      `gorm:"column:dob"`
	Email      string      `gorm:"column:email"`
	Password   string      `gorm:"column:password"`
	Phone      string      `gorm:"column:phone"`
	Address    string      `gorm:"column:address"`
	Salary     json.Number `gorm:"column:salary"`
	Position   string      `gorm:"column:position"`
	CreatedAt  time.Time   `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time   `gorm:"column:updated_at"`
}
