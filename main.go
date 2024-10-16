package main

import (
	"log"
	"net/http"
)

func main() {
	// Define HTTP routes and link them to handler functions
	http.HandleFunc("/tasks", handleTasks)     // Handles POST and GET
	http.HandleFunc("/tasks/", handleTaskByID) // Handles GET, PUT, DELETE by ID

	// Start the server on port 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
