package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	HandlerFunc           func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func ConfigureRoutes(r *mux.Router) *mux.Router {
	route := userRoutes

	for _, route := range route {
		r.HandleFunc(route.URI, route.HandlerFunc).Methods(route.Method)
	}

	return r
}
