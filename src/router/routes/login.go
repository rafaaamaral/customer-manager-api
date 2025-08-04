package routes

import (
	"customer-manager-api/src/controllers"
	"net/http"
)

var loginRoute = Route{
	URI:                   "/login",
	Method:                http.MethodPost,
	HandlerFunc:           controllers.Login,
	RequireAuthentication: false,
}
