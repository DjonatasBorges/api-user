package routes

import (
	"log"
	"net/http"

	"github.com/DjonatasBorges/api-user/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	SetupLoginRoutes(r)

	SetupUsersRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
