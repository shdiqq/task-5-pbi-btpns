package routes

import (
	"github.com/gorilla/mux"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/controllers/usercontroller"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/middleware"
)

func UserRoutes(router *mux.Router) {
	user := router.PathPrefix("/users").Subrouter()

	// Middleware
	user.Use(middleware.Auth)

	user.HandleFunc("", usercontroller.ListUser).Methods("GET")
	user.HandleFunc("/{userId}", usercontroller.UpdateUser).Methods("PUT")
	user.HandleFunc("/{userId}", usercontroller.DeleteUser).Methods("DELETE")
}
