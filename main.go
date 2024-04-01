package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID int `json:"id"`
	Text string `json:"text"`
}

var tasks []Task

func main() {
	http.HandleFunc("/tasks", handleTasks)
	http.HandleFunc("/tasks/add", addTask)
	http.HandleFunc("/tasks/delete", deleteTask)
	http.ListenAndServe(":8080", nil)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasks(w)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTasks(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	tasks = append(tasks, task)
	w.WriteHeader(http.StatusCreated)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var taskID int
	err := json.NewDecoder(r.Body).Decode(&taskID)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
