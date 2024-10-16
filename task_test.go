package main

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	title := "Test Task"
	description := "This is a test task"
	priority := 1
	task := CreateTask(title, description, priority)
	if task.Title != title || task.Description != description || task.Priority != priority {
		t.Errorf("Expected task to have title %s, description %s, and priority %d", title, description, priority)
	}
	if _, exists := tasks[task.ID]; !exists {
		t.Errorf("Expected task to be saved in tasks map")
	}
}

func TestGetTask(t *testing.T) {
	task := CreateTask("Test Task", "Test", 1)
	retrievedTask, err := GetTask(task.ID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if retrievedTask.ID != task.ID {
		t.Errorf("Expected task ID %s, got %s", task.ID, retrievedTask.ID)
	}
}

func TestGetTask_NotExist(t *testing.T) {
	_, err := GetTask("non-existent-id")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	expectedErr := "task not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %s, got %s", expectedErr, err.Error())
	}
}

func TestUpdateTask(t *testing.T) {
	task := CreateTask("Test Task", "Test", 1)
	updatedTitle := "Updated Task"
	updatedDescription := "Updated Description"
	updatedPriority := 2
	updatedTask, err := UpdateTask(task.ID, updatedTitle, updatedDescription, updatedPriority)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if updatedTask.Title != updatedTitle || updatedTask.Description != updatedDescription || updatedTask.Priority != updatedPriority {
		t.Errorf("Task not updated correctly")
	}
}

func TestUpdateTask_NotExist(t *testing.T) {
	_, err := UpdateTask("non-existent-id", "Updated Title", "Updated Description", 2)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	expectedErr := "task not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %s, got %s", expectedErr, err.Error())
	}
}

func TestDeleteTask(t *testing.T) {
	task := CreateTask("Test Task", "Test", 1)
	err := DeleteTask(task.ID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if _, exists := tasks[task.ID]; exists {
		t.Errorf("Expected task to be deleted from tasks map")
	}
}

func TestDeleteTask_NotExist(t *testing.T) {
	err := DeleteTask("non-existent-id")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	expectedErr := "task not found"
	if err.Error() != expectedErr {
		t.Errorf("Expected error message %s, got %s", expectedErr, err.Error())
	}
}
