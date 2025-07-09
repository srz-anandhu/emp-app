package dto

import (
	"encoding/json"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a *AdminLogin) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		return err
	}
	return nil
}

func (a *AdminLogin) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}

	return nil
}

type AdminToken struct {
	AdminLoginResponse
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type AdminLoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Password string `json:"password"`
	Role string `json:"role"`
}

// Employee management field for Admins
type AddEmployeeDetails struct {
	EmployeeID string      `json:"employee_id"`
	FullName   string      `json:"name"`
	Email      string      `json:"email"`
	Password   string      `json:"password"`
	Address    string      `json:"address"`
	Phone      string      `json:"phone"`
	DOB        string      `json:"dob"`
	Position   string      `json:"position"`
	Salary     json.Number `json:"salary"`
}

func (a *AddEmployeeDetails) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		return err
	}
	return nil
}

func (a *AddEmployeeDetails) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil
}

type AdminDetails struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (a *AdminDetails) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		return err
	}
	return nil
}

func (a *AdminDetails) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil
}
