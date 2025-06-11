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

	adminRepo := repository.NewAdminRepo(db)
	adminService := service.NewAdminService(adminRepo)
	adminController := controller.NewAdminController(adminService)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"} ,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.Route("/admin", func(r chi.Router) {
		r.Post("/login", adminController.Login)
		r.With(middleware.AuthMiddleware).Post("/create", adminController.AddEmployee)
	})

	r.Route("/employee", func(r chi.Router) {
	//	r.Post("/signup", empController.CreateEmployee)
		r.Post("/login", empController.Login)
		r.With(middleware.RequireAuth).Put("/{id}", empController.UpdateEmployee)
		r.With(middleware.RequireAuth).Get("/{id}", empController.GetEmployee)
		r.With(middleware.RequireAuth).Get("/", empController.GetAllEmployees)
		r.With(middleware.RequireAuth).Post("/logout", empController.Logout)
		r.With(middleware.RequireAuth).Put("/{id}/password", empController.ChangePassword)
	})

	return r
}
