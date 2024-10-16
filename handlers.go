package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// handleTasks handles both POST (create) and GET (list) requests for tasks
func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// Create a new task
		var task Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		if task.Title == "" || task.Priority == 0 {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}
		newTask := CreateTask(task.Title, task.Description, task.Priority)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newTask)

	case "GET":
		// List all tasks
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleTaskByID handles GET, PUT, DELETE for individual tasks by ID
func handleTaskByID(w http.ResponseWriter, r *http.Request) {
	// Extract the task ID from the URL
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	if id == "" {
		http.Error(w, "Task ID required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		// Retrieve a task by ID
		task, err := GetTask(id)
		if err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)

	case "PUT":
		// Update a task
		var updatedTask Task
		if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		task, err := UpdateTask(id, updatedTask.Title, updatedTask.Description, updatedTask.Priority)
		if err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)

	case "DELETE":
		// Delete a task
		if err := DeleteTask(id); err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
