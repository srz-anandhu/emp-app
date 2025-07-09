package service

import (
	"emp-app/app/domain"
	"emp-app/app/dto"
	"emp-app/app/repository"
	"emp-app/pkg/helpers/e"
	"emp-app/pkg/helpers/hash"
	jwtPackage "emp-app/pkg/helpers/jwt"
	"fmt"
	"net/http"
	"strings"
)

type AdminService interface {
	Login(r *http.Request) (*dto.AdminToken, error)
	AddEmployee(r *http.Request) (*domain.Employee, error)
	AddAdmin(r *http.Request) (*dto.AdminLoginResponse, error)
}

type AdminServiceImpl struct {
	adminRepo repository.AdminRepo
}

func NewAdminService(adminRepo repository.AdminRepo) AdminService {
	return &AdminServiceImpl{
		adminRepo: adminRepo,
	}
}

func (s *AdminServiceImpl) Login(r *http.Request) (*dto.AdminToken, error) {
	req := &dto.AdminLogin{}

	if err := req.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "request parse error", err)
	}
	if err := req.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "request validation error", err)
	}

	admin, err := s.adminRepo.FindAdminByEmail(req.Email)
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "no admin found by email", err)
	}
	// password, err := hash.HashPassword(req.Password)
	// if err != nil {
	// 	return nil, e.NewError(e.ErrInternalServer, "password hashing failed for admin", err)
	// }
	if err := hash.ComparePassword(req.Password, admin.Password); err != nil {
		return nil, e.NewError(e.ErrInternalServer, "admin password doesnt match", err)
	}
	// Generate token
	accessToken, refreshToken, err := jwtPackage.GenerateTokens(admin.ID, admin.Email, "admin")
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "could't generate token for admin", err)
	}
	// Adding fetched admin details to login reponse
	adminRes := &dto.AdminLoginResponse{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
		Role:  admin.Role,
	}

	return &dto.AdminToken{
			AdminLoginResponse: *adminRes,
			AccessToken:        accessToken,
			RefreshToken:       refreshToken,
		},
		nil
}

func (s *AdminServiceImpl) AddEmployee(r *http.Request) (*domain.Employee, error) {

	body := &dto.AddEmployeeDetails{}
	if err := body.Parse(r); err != nil {
		return nil, e.NewError(e.ErrDecodeRequestBody, "can't decode employee create request", err)
	}

	if err := body.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "can't validate employee create request", err)
	}
	
// Generate if there is no employee ID provided
		if strings.TrimSpace(body.EmployeeID) == "" {
		var count int64
		if err := s.adminRepo.CountEmployees(&count); err != nil {
			return nil, e.NewError(e.ErrInternalServer, "can't count employees", err)
		}
		body.EmployeeID = fmt.Sprintf("SRZEE%03d", count+1)
	}

	// Password hashing
	password, err := hash.HashPassword(body.Password)

	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "password hashing error", err)
	}

	// Passing hashed password to body
	body.Password = password

	// Calling repo function
	emp, err := s.adminRepo.AddEmployee(*body)

	if err != nil {
		if strings.Contains(err.Error(), "conflict") {

			return nil, e.NewError(e.ErrConflict, "Employee ID already exists", err)
		}
		return nil, e.NewError(e.ErrInternalServer, "can't create employee", err)
	}

	return emp, nil
}

func (s *AdminServiceImpl) AddAdmin(r *http.Request) (*dto.AdminLoginResponse, error) {
	body := &dto.AdminDetails{}

	if err := body.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "request body parse error ", err)
	}

	if err := body.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "request validation error ", err)
	}

	hashedPass, err := hash.HashPassword(body.Password)
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "password hashing failed", err)
	}

	body.Password = hashedPass

	result, err := s.adminRepo.AddNewAdmin(*body)
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "creating admin failed", err)
	}

	admin := &dto.AdminLoginResponse{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
		Role:  result.Role,
	}

	return admin, nil
}
