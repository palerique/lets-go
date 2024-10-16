package main

import (
	"errors"
	"fmt"
)

// Task struct represents a task with an ID, title, description, etc.
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

// In-memory task storage (map of task ID to Task)
var tasks = map[string]Task{}

// ID generator
var idCounter = 1

// CreateTask adds a new task to the tasks map
func CreateTask(title, description string, priority int) Task {
	id := fmt.Sprintf("%d", idCounter) // Create a new ID
	idCounter++
	task := Task{
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
	}
	tasks[task.ID] = task
	return task
}

// GetTask retrieves a task by its ID
func GetTask(id string) (Task, error) {
	task, exists := tasks[id]
	if !exists {
		return Task{}, errors.New("task not found")
	}
	return task, nil
}

// UpdateTask updates an existing task
func UpdateTask(id, title, description string, priority int) (Task, error) {
	task, err := GetTask(id)
	if err != nil {
		return Task{}, err
	}
	task.Title = title
	task.Description = description
	task.Priority = priority
	tasks[id] = task
	return task, nil
}

// DeleteTask removes a task by its ID
func DeleteTask(id string) error {
	_, exists := tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(tasks, id)
	return nil
}
