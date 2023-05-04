package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Msmpunk/go-server/core"
	"github.com/Msmpunk/go-server/models"
)

var db = core.Database()

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prepareInsertion, err := db.Prepare("INSERT INTO tasks (task, description) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = prepareInsertion.Exec(task.Task, task.Description)
	if err != nil {
		fmt.Println(err.Error())
	}
	json.NewEncoder(w).Encode("Task created.")
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.Task
	jErr := json.NewDecoder(r.Body).Decode(&task)
	if jErr != nil {
		panic(jErr)
	}

	_, err := db.Query("UPDATE tasks SET task = ? WHERE id = ?", task.Task, task.Id)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode("Task updated.")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var results []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Task, &task.Description)
		if err != nil {
			panic(err)
		}
		results = append(results, task)
	}

	json.NewEncoder(w).Encode(results)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		panic(err)
	}

	_, errx := db.Query("DELETE FROM tasks WHERE id = ?", task.Id)
	if errx != nil {
		panic(errx.Error())
	}

	json.NewEncoder(w).Encode("Task deleted.")
}
