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

func (c *EmployeeControllerImpl) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	result, err := c.empService.CreateEmployee(r)
	if err != nil {
		httpErr := e.NewApiError(err, "can't create employee")
		response.Fail(w, http.StatusInternalServerError, httpErr.Code, httpErr.Message, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, result)
}