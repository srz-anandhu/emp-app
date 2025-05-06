package service

import (
	"emp-app/app/domain"
	"emp-app/app/dto"
	"emp-app/app/repository"
	"net/http"
)

type EmployeeService interface {
	CreateEmployee(r *http.Request) (*domain.Employee, error)
	GetEmployee(r *http.Request) (*domain.Employee, error)
	UpdateEmployee(r *http.Request) error
	GetAllEmployees(r *http.Request) ([]*domain.Employee, error)
}

type EmployeeServiceImpl struct {
	empRepo repository.EmployeeRepo
}

// Constructor function
func NewEmployeeService(empRepo repository.EmployeeRepo) EmployeeService {
	return &EmployeeServiceImpl{
		empRepo: empRepo,
	}
}

func (s *EmployeeServiceImpl) CreateEmployee(r *http.Request) (*domain.Employee, error) {

	body := &dto.EmployeeCreateRequest{}
	if err := body.Parse(r); err != nil {
		return nil, err
	}

	if err := body.Validate(); err != nil {
		return nil, err
	}
	emp, err := s.empRepo.CreateEmployee(body)

	if err != nil {
		return nil, err
	}

	return emp, nil
}

func (s *EmployeeServiceImpl) GetEmployee(r *http.Request) (*domain.Employee, error) {
	req := &dto.EmployeeRequest{}
	if err := req.Parse(r); err != nil {
		return nil, err
	}

	if err := req.Validate(r); err != nil {
		return nil, err
	}

	emp, err := s.empRepo.GetEmployee(req)
	if err != nil {
		return nil, err
	}

	var employee domain.Employee

	employee.ID = emp.ID
	employee.Name = emp.Name
	employee.Email = emp.Email
	employee.Address = emp.Address
	employee.DOB = emp.DOB
	employee.Phone = emp.Phone
	employee.Position = emp.Position
	employee.Salary = emp.Salary
	employee.CreatedAt = emp.CreatedAt
	employee.UpdatedAt = emp.UpdatedAt

	return &employee, nil
}

func (s *EmployeeServiceImpl) UpdateEmployee(r *http.Request) error {

	body := &dto.EmployeeUpdateRequest{}

	if err := body.Parse(r); err != nil {
		return err
	}

	if err := body.Validate(); err != nil {
		return err
	}

	if err := s.empRepo.UpdateEmployee(body); err != nil {
		return err
	}

	return nil

}

func (s *EmployeeServiceImpl) GetAllEmployees(r *http.Request) ([]*domain.Employee, error) {
	results, err := s.empRepo.GetAllEmployees()

	if err != nil {
		return nil, err
	}

	var employees []*domain.Employee

	for _, val := range results {

		var emp domain.Employee

		emp.ID = val.ID
		emp.Name = val.Name
		emp.Email = val.Email
		emp.Address = val.Address
		emp.DOB = val.DOB
		emp.Phone = val.Phone
		emp.Position = val.Position
		emp.Salary = val.Salary
		emp.CreatedAt = val.CreatedAt
		emp.UpdatedAt = val.UpdatedAt

		employees = append(employees, &emp)

	}

	return employees, nil

}
