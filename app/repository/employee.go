package repository

import (
	"emp-app/app/domain"
	"emp-app/app/dto"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type EmployeeRepo interface {
	CreateEmployee(createReq *dto.EmployeeCreateRequest) (*domain.Employee, error)
	GetEmployee(empReq *dto.EmployeeRequest) (*domain.Employee, error)
	UpdateEmployee(empUpdateReq *dto.EmployeeUpdateRequest) error
	GetAllEmployees() ([]*domain.Employee, error)
}

type EmployeeRepoImpl struct {
	db *gorm.DB
}

// Constructor funciton
func NewEmployeeRepo(db *gorm.DB) EmployeeRepo {
	return &EmployeeRepoImpl{
		db: db,
	}
}

// For checking implementation of EmployeeRepo interface
var _ EmployeeRepo = (*EmployeeRepoImpl)(nil)

func (r *EmployeeRepoImpl) CreateEmployee(createReq *dto.EmployeeCreateRequest) (*domain.Employee, error) {
	employee := &domain.Employee{
		Name:     createReq.Name,
		DOB:      createReq.DOB,
		Email:    createReq.Email,
		Password: createReq.Password,
		Phone:    createReq.Phone,
		Address:  createReq.Address,
		Salary:   createReq.Salary,
		Position: createReq.Position,
	}

	result := r.db.Create(employee)
	if result.Error != nil {
		return nil, result.Error
	}

	return employee, nil
}

func (r *EmployeeRepoImpl) GetEmployee(empReq *dto.EmployeeRequest) (*domain.Employee, error) {
	emp := &domain.Employee{}

	result := r.db.Where("id = ?", empReq.ID).First(emp)
	if result.Error != nil {
		// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 	return nil, fmt.Errorf("resource not found")
		// }
		return nil, result.Error
	}
	return emp, nil
}

func (r *EmployeeRepoImpl) UpdateEmployee(empUpdateReq *dto.EmployeeUpdateRequest) error {
	result := r.db.Table("employees").Where("id = ?", empUpdateReq.ID).Updates(map[string]any{

		"name":       empUpdateReq.Name,
		"dob":        empUpdateReq.DOB,
		"email":      empUpdateReq.Email,
		"password":   empUpdateReq.Password,
		"phone":      empUpdateReq.Phone,
		"address":    empUpdateReq.Address,
		"position":   empUpdateReq.Position,
		"salary":     empUpdateReq.Salary,
		"updated_at": time.Now().UTC(),
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no employee found with ID : %d to update", empUpdateReq.ID)
	}

	return nil
}

func (r *EmployeeRepoImpl) GetAllEmployees() ([]*domain.Employee, error) {
	var employees []*domain.Employee
	result := r.db.Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}

	return employees, nil
}
