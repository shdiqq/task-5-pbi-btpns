package routes

import (
	"github.com/gorilla/mux"
)

func RoutesIndex(router *mux.Router) {
	AuthRoutes(router)
	UserRoutes(router)
	PhotoRoutes(router)
}
