package repository

import (
	"emp-app/app/domain"
	"emp-app/app/dto"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type EmployeeRepo interface {
	//CreateEmployee(createReq *dto.EmployeeCreateRequest) (*domain.Employee, error)
	GetEmployee(empReq *dto.EmployeeRequest) (*domain.Employee, error)
	UpdateEmployee(empUpdateReq *dto.EmployeeUpdateRequest) error
	GetAllEmployees() ([]*domain.Employee, error)
	FindUserByEmail(email string) (*domain.Employee, error)
	GetPasswordFromID(empPassReq *dto.EmployeePassRequest) (string, error)
	ChangePassword(empPassChange *dto.EmployeePassChange) error
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

func (r *EmployeeRepoImpl) FindUserByEmail(email string) (*domain.Employee, error) {
	emp := &domain.Employee{}
	// Get underlying sql.DB for connection management
	sqlDB, err := r.db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	// Verify connection is alive
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}
	result := r.db.Where("email= ?", email).First(emp)
	if result.Error != nil {
		return nil, result.Error
	}
	return emp, nil
}

// func (r *EmployeeRepoImpl) CreateEmployee(createReq *dto.EmployeeCreateRequest) (*domain.Employee, error) {
// 	employee := &domain.Employee{
// 		FullName:     createReq.Name,
// 		DOB:      createReq.DOB,
// 		Email:    createReq.Email,
// 		Password: createReq.Password,
// 		Phone:    createReq.Phone,
// 		Address:  createReq.Address,
// 		Salary:   createReq.Salary,
// 		Position: createReq.Position,
// 	}

// 	result := r.db.Create(employee)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return employee, nil
// }

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

		"name":       empUpdateReq.FullName,
		"dob":        empUpdateReq.DOB,
		"email":      empUpdateReq.Email,
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

// Get password from employee ID
func (r *EmployeeRepoImpl) GetPasswordFromID(empReq *dto.EmployeePassRequest) (string, error) {
	empPass := &dto.EmployeePassRequest{}
	result := r.db.Table("employees").Select("password").Where("id = ?", empReq.ID).First(empPass)
	if result.Error != nil {
		return "", result.Error
	}
	return empPass.Password, nil
}

func (r *EmployeeRepoImpl) ChangePassword(empPassChange *dto.EmployeePassChange) error {
	if empPassChange.NewPassword == nil {
		return fmt.Errorf("new password must not be nil")
	}
	result := r.db.Table("employees").Where("id =?", empPassChange.ID).Update("password", *empPassChange.NewPassword)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no employee found with ID")
	}
	return nil
}
