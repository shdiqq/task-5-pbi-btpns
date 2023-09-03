package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/config"
	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/routes"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()

	router := mux.NewRouter()
	routes.RoutesIndex(router)

	log.Println("[APP] Server is listening to port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), router)
}
