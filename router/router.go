package router

import (
	"extensao-api/router/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	rotas := mux.NewRouter()
	routes.Register(rotas)
	return rotas
}
