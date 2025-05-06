package service

import (
	"emp-app/app/domain"
	"emp-app/app/dto"
	"emp-app/app/repository"
	"emp-app/pkg/helpers/e"
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
		return nil, e.NewError(e.ErrDecodeRequestBody, "can't decode employee create request", err)
	}

	if err := body.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "can't validate employee create request", err)
	}
	emp, err := s.empRepo.CreateEmployee(body)

	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "cant create employee", err)
	}

	return emp, nil
}

func (s *EmployeeServiceImpl) GetEmployee(r *http.Request) (*domain.Employee, error) {
	req := &dto.EmployeeRequest{}
	if err := req.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "employee request parse error", err)
	}

	if err := req.Validate(r); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "employee request validation error", err)
	}

	emp, err := s.empRepo.GetEmployee(req)
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "getting employee failed", err)
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
		return e.NewError(e.ErrDecodeRequestBody, "can't decode employee update request", err)
	}

	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "can't validate employee update request", err)
	}

	if err := s.empRepo.UpdateEmployee(body); err != nil {
		return e.NewError(e.ErrInternalServer, "update employee failed", err)
	}

	return nil

}

func (s *EmployeeServiceImpl) GetAllEmployees(r *http.Request) ([]*domain.Employee, error) {
	results, err := s.empRepo.GetAllEmployees()

	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "can't get all employees", err)
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
