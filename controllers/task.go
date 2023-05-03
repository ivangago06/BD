package controllers

import (
	"BD/core"
	"BD/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var db = core.Database()

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.Task
	err := json.NewEncoder(r.Body).Encode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prepareIn, err := db.Prepare("INSERT INTO task (task, descripcion) VALUES (?, ?)")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = prepareIn.Exec(task.Task, task.Descripcion)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.NewEncoder(w).Encode("La tarea fue creada")

}

func UpdateTask() {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.Task

	err := json.NewEncoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Query("UPDATE task SET task=? WHERE id=?", task.Task, task.Id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	json.NewEncoder(w).Encode("La tarea fue actualizada")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	rows, err := db.Query("SELECT * FROM TASK")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var result []models.Task

	for rows.Next() {

		var task models.Task

		err := rows.Scan(&task.Id, task.Task, task.Descripcion)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, task)

	}

	json.NewEncoder(w).Encode(result)

}
