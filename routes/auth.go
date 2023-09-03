package routes

import (
	"github.com/gorilla/mux"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/controllers/authcontroller"
)

func AuthRoutes(router *mux.Router) {
	auth := router.PathPrefix("/users").Subrouter()

	auth.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	auth.HandleFunc("/register", authcontroller.Register).Methods("POST")
	auth.HandleFunc("/login", authcontroller.Login).Methods("POST")
}
