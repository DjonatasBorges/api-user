package routes

import (
	"github.com/DjonatasBorges/api-user/controllers"
	"github.com/gorilla/mux"
)

func SetupLoginRoutes(r *mux.Router) {
	r.HandleFunc("/login", controllers.Login).Methods("POST")
}
