package service

import (
	"emp-app/app/dto"
	"emp-app/app/repository"
	"emp-app/pkg/helpers/e"
	"emp-app/pkg/helpers/hash"
	jwtPackage "emp-app/pkg/helpers/jwt"
	"net/http"
)

type AdminService interface {
	Login(r *http.Request) (*dto.AdminToken, error)
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
	password, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "password hashing failed for admin", err)
	}
	if err := hash.ComparePassword(password, admin.Password); err != nil {
		return nil, e.NewError(e.ErrInternalServer, "admin password doesnt match", err)
	}
	// Generate token
	accessToken, refreshToken, err := jwtPackage.GenerateTokens(admin.ID, admin.Email, "admin")
	if err != nil {
		return nil, e.NewError(e.ErrInternalServer, "could't generate token for admin", err)
	}
	// Adding fetched admin details to login reponse
	adminRes := &dto.AdminLoginResponse{
		ID: admin.ID,
		Name: admin.Name,
		Email: admin.Email,
		Role: admin.Role,
	}

	return &dto.AdminToken{
		AdminLoginResponse: *adminRes,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	},
	nil
}