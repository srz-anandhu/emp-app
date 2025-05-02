package dto

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	validator "github.com/go-playground/validator/v10"
)

type EmployeeCreateRequest struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	DOB      time.Time `json:"dob"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Salary   float64   `json:"salary"`
	Position string    `json:"position"`
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

type EmployeeLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (e *EmployeeLogin) Validate() error {
	validate := validator.New()
	if err := validate.Struct(e); err != nil {
		return err
	}
	return nil
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
type EmployeeResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	DOB       time.Time `json:"dob"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Salary    float64   `json:"salary"`
	Position  string    `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EmployeeUpdateRequest struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	DOB      time.Time `json:"dob"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Salary   float64   `json:"salary"`
	Position string    `json:"position"`
}
