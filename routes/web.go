package routes

import (
	"BD/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("api/v1/get-file", controllers.GetDoc).Methods("GET")
	router.HandleFunc("api/v1/task", controllers.CreateTask).Methods("POST")
	router.HandleFunc("api/v1/task", controllers.GetAllTasks).Methods("GET")
	router.HandleFunc("api/v1/task-update", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("api/v1/task", controllers.ShouUsers).Methods("GET")
	return router
}
