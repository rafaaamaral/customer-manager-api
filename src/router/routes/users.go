package routes

import (
	"customer-manager-api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		HandlerFunc:           controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		HandlerFunc:           controllers.GetUsers,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodGet,
		HandlerFunc:           controllers.GetUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodPut,
		HandlerFunc:           controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{id}",
		Method:                http.MethodDelete,
		HandlerFunc:           controllers.DeleteUser,
		RequireAuthentication: true,
	},
}
