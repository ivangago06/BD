package routes

import (
	"BD/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("api/v1/get-file", controllers.GetDoc).Methods("GET")

	return router
}
