package dto

import (
	"emp-app/app/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	validator "github.com/go-playground/validator/v10"
)

type EmployeeCreateRequest struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	DOB      string      `json:"dob"`
	Email    string      `json:"email" validate:"required"`
	Password string      `json:"password" validate:"required"`
	Phone    string      `json:"phone"`
	Address  string      `json:"address"`
	Salary   json.Number `json:"salary"`
	Position string      `json:"position"`
}

func (e *EmployeeCreateRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		return err
	}
	return nil
}

func (e *EmployeeCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		return err
	}
	return nil
}

// Token
type Token struct {
	EmployeeResp domain.Employee
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type EmployeeLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (e *EmployeeLogin) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		return err
	}

	return nil
}

func (e *EmployeeLogin) Validate() error {
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		return err
	}
	return nil
}

type LoginToken struct {
	EmpResp      EmployeeLoginResp
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type EmployeeLoginResp struct {
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Dob      string      `json:"dob"`
	Address  string      `json:"address"`
	Phone    string      `json:"phone"`
	Position string      `json:"position"`
	Salary   json.Number `json:"salary"`
}

type EmployeeRequest struct {
	ID int `validate:"required"`
}

func (e *EmployeeRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	e.ID = intID
	return nil
}

func (e *EmployeeRequest) Validate(r *http.Request) error {
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		return err
	}
	return nil
}

// After employee creation this response will sent to front-end
// type EmployeeResponse struct {
// 	ID        int       `json:"id"`
// 	Name      string    `json:"name"`
// 	DOB       string    `json:"dob"`
// 	Email     string    `json:"email"`
// 	Phone     string    `json:"phone"`
// 	Address   string    `json:"address"`
// 	Salary    float64   `json:"salary"`
// 	Position  string    `json:"position"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

type EmployeeUpdateRequest struct {
	ID       int      `json:"id"`
	Name     *string  `json:"name"`
	DOB      *string  `json:"dob"`
	Email    *string  `json:"email"`
	Password *string  `json:"password"`
	Phone    *string  `json:"phone"`
	Address  *string  `json:"address"`
	Salary   *float64 `json:"salary"`
	Position *string  `json:"position"`
}

func (e *EmployeeUpdateRequest) Parse(r *http.Request) error {
	// Get ID from Request
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	e.ID = intID

	// Decode to EmployeeUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(e); err != nil {
		return err
	}

	return nil
}

func (e *EmployeeUpdateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		return err
	}
	return nil
}
