package routes

import (
	"github.com/DjonatasBorges/api-user/controllers"
	"github.com/DjonatasBorges/api-user/middleware"
	"github.com/gorilla/mux"
)

func SetupUsersRoutes(r *mux.Router) {

	r.HandleFunc("/users", controllers.PostUser).Methods("POST")

	protectedRoutes := r.PathPrefix("/users").Subrouter()

	protectedRoutes.Use(middleware.AuthenticationMiddleware)

	protectedRoutes.HandleFunc("", controllers.GetAllUsers).Methods("GET")
	protectedRoutes.HandleFunc("/{id}", controllers.GetUserById).Methods("GET")
	protectedRoutes.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE")
}
