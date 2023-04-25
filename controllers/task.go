package controllers

import (
	"BD/core"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var db = core.Database()

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	if params["task"] == "" {
		return
	}

	prepareIn, err := db.Prepare("INSET INTO task (task, name, adress) VALUES ('?', 0)")

	if err != nil {
		log.Fatal("Error al insertar en la BD")
	}

	prepareIn.Exec(params["task"])
	json.NewEncoder(w).Encode("Tarea Creada")

}
