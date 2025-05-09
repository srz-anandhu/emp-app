package controller

import (
	"emp-app/app/service"
	"emp-app/pkg/helpers/e"
	"emp-app/pkg/response"
	"net/http"
)

type EmployeeController interface {
	CreateEmployee(w http.ResponseWriter, r *http.Request)
	GetEmployee(w http.ResponseWriter, r *http.Request)
	UpdateEmployee(w http.ResponseWriter, r *http.Request)
	GetAllEmployees(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type EmployeeControllerImpl struct {
	empService service.EmployeeService
}

// Constructor function
func NewEmployeeController(empService service.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{
		empService: empService,
	}
}

func (c *EmployeeControllerImpl) Login(w http.ResponseWriter, r *http.Request) {

	result, err := c.empService.Login(r)
	if err != nil {
		httpErr := e.NewApiError(err, "can't login")
		response.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	response.Success(w, http.StatusOK, result)
}

func (c *EmployeeControllerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	if err := c.empService.Logout(r); err != nil {
		httpErr := e.NewApiError(err, "can't blacklist token")
		response.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "logout successfully...")
}

func (c *EmployeeControllerImpl) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	result, err := c.empService.CreateEmployee(r)
	if err != nil {
		httpErr := e.NewApiError(err, "can't create employee")
		response.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, result)
}

func (c *EmployeeControllerImpl) GetEmployee(w http.ResponseWriter, r *http.Request) {
	result, err := c.empService.GetEmployee(r)
	if err != nil {
		httpErr := e.NewApiError(err, "can't get employee")
		response.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}
	response.Success(w, http.StatusOK, result)
}

func (c *EmployeeControllerImpl) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	if err := c.empService.UpdateEmployee(r); err != nil {
		httpErr := e.NewApiError(err, "can't update employee")
		response.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "updated employee successfully")
}

func (c *EmployeeControllerImpl) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	results, err := c.empService.GetAllEmployees(r)
	if err != nil {
		httpErr := e.NewApiError(err, "can't get all employees")
		response.Fail(w, httpErr.StatusCode, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	response.Success(w, http.StatusOK, results)
}
