package repository

import (
	"emp-app/app/domain"
	"emp-app/app/dto"
	"emp-app/pkg/helpers/e"

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
