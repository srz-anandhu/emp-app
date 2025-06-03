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
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
