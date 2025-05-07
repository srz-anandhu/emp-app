package app

import (
	"emp-app/app/controller"
	"emp-app/app/repository"
	"emp-app/app/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)
func ApiRoute(db *gorm.DB) chi.Router {

	// Dependency injection
	empRepo := repository.NewEmployeeRepo(db)
	empService := service.NewEmployeeService(empRepo)
	empController := controller.NewEmployeeController(empService)

	r := chi.NewRouter()

	r.Route("/employee", func(r chi.Router) {
		r.Post("/signup", empController.CreateEmployee)
		r.Put("/{id}", empController.UpdateEmployee)
		r.Get("/{id}", empController.GetEmployee)
		r.Get("/", empController.GetAllEmployees)
		r.Post("/login", empController.Login)
	})

	return r
}