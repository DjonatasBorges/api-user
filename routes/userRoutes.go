package routes

import (
	"log"
	"net/http"

	"github.com/DjonatasBorges/api-user/controllers"
	"github.com/DjonatasBorges/api-user/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/users", controllers.GetAllUsers).Methods("Get")
	r.HandleFunc("/users/{id}", controllers.GetUserById).Methods("Get")
	r.HandleFunc("/users", controllers.PostUser).Methods("Post")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("Delete")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
