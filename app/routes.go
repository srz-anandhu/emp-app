package app

import (
	"emp-app/app/controller"
	"emp-app/app/repository"
	"emp-app/app/service"
	"emp-app/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
)

func ApiRoute(db *gorm.DB) chi.Router {

	// Dependency injection
	empRepo := repository.NewEmployeeRepo(db)
	empService := service.NewEmployeeService(empRepo)
	empController := controller.NewEmployeeController(empService)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"} ,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.Route("/employee", func(r chi.Router) {
		r.Post("/signup", empController.CreateEmployee)
		r.Post("/login", empController.Login)
		r.With(middleware.AuthMiddleware).Put("/{id}", empController.UpdateEmployee)
		r.With(middleware.AuthMiddleware).Get("/{id}", empController.GetEmployee)
		r.With(middleware.AuthMiddleware).Get("/", empController.GetAllEmployees)
		r.With(middleware.AuthMiddleware).Post("/logout", empController.Logout)
		r.With(middleware.AuthMiddleware).Put("/{id}/password", empController.ChangePassword)
	})

	return r
}
