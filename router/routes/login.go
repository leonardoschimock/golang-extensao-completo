package routes

import (
	"extensao-api/controller"
	"net/http"
)

var loginRoutes = []Routes{
	{
		URI:    "/login",
		Method: http.MethodPost,
		Func:   controller.Login,
		Auth:   false,
	},
}
