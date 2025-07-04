package repository

import (
	"emp-app/app/domain"
	"emp-app/app/dto"
	"emp-app/pkg/helpers/e"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AdminRepo interface {
	FindAdminByEmail(email string) (*domain.Admin, error)
	AddEmployee(empDetails dto.AddEmployeeDetails) (*domain.Employee, error)
	AddNewAdmin(adminDetails dto.AdminDetails) (*domain.Admin, error)
}

type AdminRepoImpl struct {
	db *gorm.DB
}

func NewAdminRepo(db *gorm.DB) AdminRepo {
	return &AdminRepoImpl{
		db: db,
	}
}

// For checking the implementation of AdminRepo interface
var _ AdminRepo = (*AdminRepoImpl)(nil)

func (r *AdminRepoImpl) FindAdminByEmail(email string) (*domain.Admin, error) {
	admin := &domain.Admin{}
	result := r.db.Table("admins").Where("email = ?", email).First(admin)
	if result.Error != nil {
		return nil, e.NewError(e.ErrResourceNotFound, "no admin found with given email", result.Error)
	}
	return admin, nil
}

// Add employee
func (r *AdminRepoImpl) AddEmployee(empDetails dto.AddEmployeeDetails) (*domain.Employee, error) {

	// Check employee ID already exist or not
	var existedID domain.Employee

	err := r.db.Where("employee_id = ?", empDetails.EmployeeID).First(&existedID).Error
	if err == nil {
		// record found
		return nil, fmt.Errorf("conflict : Employee ID already exist")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// database error
		return nil, err
	}
	// create employee
	employee := &domain.Employee{
		EmployeeID: empDetails.EmployeeID,
		FullName:   empDetails.FullName,
		DOB:        empDetails.DOB,
		Email:      empDetails.Email,
		Password:   empDetails.Password,
		Phone:      empDetails.Phone,
		Address:    empDetails.Address,
		Salary:     empDetails.Salary,
		Position:   empDetails.Position,
	}

	result := r.db.Create(employee)
	if result.Error != nil {
		return nil, result.Error
	}

	return employee, nil
}

func (r *AdminRepoImpl) AddNewAdmin(adminDetails dto.AdminDetails) (*domain.Admin, error) {
	admin := &domain.Admin{
		Name:     adminDetails.Name,
		Email:    adminDetails.Email,
		Password: adminDetails.Password,
		Role:     adminDetails.Role,
	}

	result := r.db.Create(admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}
