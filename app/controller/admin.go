package controller

import (
	"emp-app/app/service"
	"emp-app/pkg/helpers/e"
	"emp-app/pkg/response"
	"net/http"
)

type AdminController interface {
	Login(w http.ResponseWriter, r *http.Request)
	AddEmployee(w http.ResponseWriter, r *http.Request)
	AddNewAdmin(w http.ResponseWriter, r *http.Request)
}

type AdminControllerImpl struct {
	adminService service.AdminService
}

func NewAdminController(adminService service.AdminService) AdminController {
	return &AdminControllerImpl{
		adminService: adminService,
	}
}

func (c *AdminControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	resp, err := c.adminService.Login(r)
	if err != nil {
		httpErr := e.NewApiError(err, "admin login failed")
		response.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	response.Success(w, http.StatusOK, resp)
}

func (c *AdminControllerImpl) AddEmployee(w http.ResponseWriter, r *http.Request) {
	resp, err := c.adminService.AddEmployee(r)
	if err != nil {
		httpErr := e.NewApiError(err, "can't add employee")
		response.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	response.Success(w, http.StatusOK, resp)
}

func (c *AdminControllerImpl) AddNewAdmin(w http.ResponseWriter, r *http.Request) {
	resp, err := c.adminService.AddAdmin(r)
	if err != nil {
		httpErr := e.NewApiError(err, "can't add admin")
		response.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, resp)
}
