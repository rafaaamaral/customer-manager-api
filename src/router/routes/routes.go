package routes

import (
	"customer-manager-api/src/middlewares"
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
	route = append(route, loginRoute)

	for _, route := range route {
		if route.RequireAuthentication {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Authorize(route.HandlerFunc)),
			).Methods(route.Method)
		} else {

			r.HandleFunc(
				route.URI,
				middlewares.Logger(route.HandlerFunc),
			).Methods(route.Method)
		}
	}

	return r
}
