package router

import (
	"customer-manager-api/src/router/routes"

	"github.com/gorilla/mux"
)

// SetupRouter initializes and returns a new Gorilla Mux router instance.
// This router can be used to register routes and handle HTTP requests for the application.
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	return routes.ConfigureRoutes(r)
}
