package service

import (
	"emp-app/app/domain"
	"emp-app/app/dto"
	"emp-app/app/repository"
	"emp-app/pkg/helpers/e"
	"emp-app/pkg/helpers/hash"
	jwtpackage "emp-app/pkg/helpers/jwt"
	"net/http"
)

type EmployeeService interface {
	//CreateEmployee(r *http.Request) (*dto.Token, error)
	GetEmployee(r *http.Request) (*domain.Employee, error)
	UpdateEmployee(r *http.Request) error
	GetAllEmployees(r *http.Request) ([]*domain.Employee, error)
	Login(r *http.Request) (*dto.LoginToken, error)
	Logout(r *http.Request) error
	ChangePassword(r *http.Request) error
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

func (s *EmployeeServiceImpl) Logout(r *http.Request) error {
	// Extract token from request header
	token := jwtpackage.ExtractTokenFromHeader(r)
	if err := jwtpackage.BlackListToken(token); err != nil {
		return e.NewError(e.ErrInternalServer, "token blacklisting failed", err)
	}
	return nil
}

func (s *EmployeeServiceImpl) Login(r *http.Request) (*dto.LoginToken, error) {

	body := &dto.EmployeeLogin{}

	if err := body.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "login request parse error", err)
	}

	if err := body.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "validate error in login request", err)
	}

	// Getting existing user by email
	emp, err := s.empRepo.FindUserByEmail(body.Email)
	if err != nil {
		return nil, e.NewError(e.ErrResourceNotFound, "user not found with given email.. please signup", err)
	}

	if err := hash.ComparePassword(body.Password, emp.Password); err != nil {
		return nil, e.NewError(e.ErrInternalServer, "password doesn't match", err)
	}
	accessToken, refreshToken, err := jwtpackage.GenerateTokens(emp.ID, emp.Email, "employee")
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "token generation error", err)
	}

	empRes := &dto.EmployeeLoginResp{
		ID:       emp.ID,
		Name:     emp.FullName,
		Email:    emp.Email,
		Password: emp.Password,
		Address:  emp.Address,
		Dob:      emp.DOB,
		Phone:    emp.Phone,
		Position: emp.Position,
		Salary:   emp.Salary,
	}

	return &dto.LoginToken{
		EmpResp:      *empRes,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// func (s *EmployeeServiceImpl) CreateEmployee(r *http.Request) (*dto.Token, error) {

// 	body := &dto.EmployeeCreateRequest{}
// 	if err := body.Parse(r); err != nil {
// 		return nil, e.NewError(e.ErrDecodeRequestBody, "can't decode employee create request", err)
// 	}

// 	if err := body.Validate(); err != nil {
// 		return nil, e.NewError(e.ErrValidateRequest, "can't validate employee create request", err)
// 	}

// 	// Password hashing
// 	password, err := hash.HashPassword(body.Password)

// 	if err != nil {
// 		return nil, e.NewError(e.ErrInternalServer, "password hashing error", err)
// 	}

// 	// Passing hashed password to body
// 	body.Password = password

// 	accessToken, refresToken, err := jwtpackage.GenerateTokens(body.ID, body.Email, "employee")
// 	if err != nil {
// 		return nil, e.NewError(e.ErrInternalServer, "token generation error", err)
// 	}

// 	// Calling repo function
// 	emp, err := s.empRepo.CreateEmployee(body)

// 	if err != nil {
// 		return nil, e.NewError(e.ErrInternalServer, "cant create employee", err)
// 	}

// 	return &dto.Token{
// 		EmployeeResp: *emp,
// 		AccessToken:  accessToken,
// 		RefreshToken: refresToken,
// 	}, nil
// }

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
	employee.FullName = emp.FullName
	employee.Email = emp.Email
	employee.Password = emp.Password
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
		emp.EmployeeID = val.EmployeeID
		emp.FullName = val.FullName
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

func (s *EmployeeServiceImpl) ChangePassword(r *http.Request) error {

	passChangeReq := &dto.EmployeePassChange{}

	if err := passChangeReq.Parse(r); err != nil {
		return e.NewError(e.ErrInvalidRequest, "password change req body parse error", err)
	}
	if err := passChangeReq.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "validation error", err)
	}

	// Get the current hashed password from DB
	password, err := s.empRepo.GetPasswordFromID(&dto.EmployeePassRequest{ID: passChangeReq.ID})
	if err != nil {
		return e.NewError(e.ErrInternalServer, "password getting error", err)
	}

	// Compare current password
	if err := hash.ComparePassword(passChangeReq.CurrentPassword, password); err != nil {
		return e.NewError(e.ErrInternalServer, "password doesn't match", err)
	}
	
	// Ensure new password and confirm password match
	if passChangeReq.NewPassword == nil || passChangeReq.ConfirmPassword == nil || 
		*passChangeReq.NewPassword != *passChangeReq.ConfirmPassword {
		return e.NewError(e.ErrInvalidRequest, "new and confirm password do not match", nil)
		}
	

	// Hash new password
	hashedNewPass, err := hash.HashPassword(*passChangeReq.NewPassword)
	if err != nil {
		return e.NewError(e.ErrInternalServer, "failed to hash new password", err)
	}

	// Set hashed password and call repository to update
	passChangeReq.NewPassword = &hashedNewPass
	
	if err := s.empRepo.ChangePassword(passChangeReq); err != nil {
		return e.NewError(e.ErrInternalServer, "password changing failed", err)
	}

	return nil
}
