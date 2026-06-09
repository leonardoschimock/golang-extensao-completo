package routes

import (
	"extensao-api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
	Auth   bool
}

func Register(r *mux.Router) {
	var routes []Routes
	routes = append(routes, booksRoutes...)
	routes = append(routes, usersRoutes...)
	routes = append(routes, loginRoutes...)

	for _, route := range routes {

		handler := middleware.Logger(route.Func)

		if route.Auth {
			handler = middleware.Logger(
				middleware.Authenticate(route.Func),
			)
		}

		r.HandleFunc(
			route.URI,
			handler,
		).Methods(route.Method)
	}
}
