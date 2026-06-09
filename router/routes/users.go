package routes

import (
	"extensao-api/controller"
	"net/http"
)

var usersRoutes = []Routes{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func:   controller.CreateUser,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func:   controller.FetchUser,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodPut,
		Func:   controller.UpdateUser,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodDelete,
		Func:   controller.DeleteUser,
	},
}
