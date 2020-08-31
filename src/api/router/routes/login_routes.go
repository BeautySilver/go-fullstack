package routes

import (
	"../../controllers"
	"net/http"
)

var loginRoutes =[]Route{
	{
		Uri:     "/login",
		Method: http.MethodPost ,
		Handler: controllers.Login,
		AuthRequired: false,
	},
}
