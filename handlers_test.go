package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// setup initializes the global tasks map before each test
func setup() {
	tasks = make(map[string]Task)
}

func TestHandleTasksInvalidPayload(t *testing.T) {
	setup()

	req, err := http.NewRequest("POST", "/tasks", bytes.NewBufferString("invalid payload"))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTasks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestHandleTasksMissingFields(t *testing.T) {
	setup()

	task := Task{Title: "", Priority: 0}
	taskJSON, _ := json.Marshal(task)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTasks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestHandleTasksValidInput(t *testing.T) {
	setup()

	task := Task{Title: "New Task", Description: "New Description", Priority: 1}
	taskJSON, _ := json.Marshal(task)
	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTasks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHandleTasksMethodNotAllowed(t *testing.T) {
	setup()

	req, err := http.NewRequest("PATCH", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTasks)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestHandleTaskByIDEmptyID(t *testing.T) {
	setup()

	req, err := http.NewRequest("GET", "/tasks/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTaskByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestHandleTaskByIDGetTaskNotFound(t *testing.T) {
	setup()

	req, err := http.NewRequest("GET", "/tasks/invalid-id", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTaskByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestHandleTaskByIDGetTaskSuccess(t *testing.T) {
	setup()

	task := Task{ID: "1", Title: "Test Task", Description: "Test", Priority: 1}
	tasks[task.ID] = task

	req, err := http.NewRequest("GET", "/tasks/"+task.ID, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTaskByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHandleTaskByIDPutInvalidBody(t *testing.T) {
	setup()

	task := Task{ID: "1", Title: "Test Task", Description: "Test", Priority: 1}
	tasks[task.ID] = task

	req, err := http.NewRequest("PUT", "/tasks/"+task.ID, bytes.NewBufferString("invalid payload"))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTaskByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestHandleTaskByIDPutUpdateTaskNotFound(t *testing.T) {
	setup()

	task := Task{ID: "1", Title: "Test Task", Description: "Test", Priority: 1}
	taskJSON, _ := json.Marshal(task)
	req, err := http.NewRequest("PUT", "/tasks/invalid-id", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTaskByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestHandleTaskByIDPutUpdateTaskSuccess(t *testing.T) {
	setup()

	task := Task{ID: "1", Title: "Test Task", Description: "Test", Priority: 1}
	tasks[task.ID] = task

	updatedTask := Task{Title: "Updated Task", Description: "Updated Description", Priority: 2}
	updatedTaskJSON, _ := json.Marshal(updatedTask)
	req, err := http.NewRequest("PUT", "/tasks/"+task.ID, bytes.NewBuffer(updatedTaskJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTaskByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHandleTaskByIDDeleteTaskNotFound(t *testing.T) {
	setup()

	req, err := http.NewRequest("DELETE", "/tasks/invalid-id", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTaskByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestHandleTaskByIDDeleteTaskSuccess(t *testing.T) {
	setup()

	task := Task{ID: "1", Title: "Test Task", Description: "Test", Priority: 1}
	tasks[task.ID] = task

	req, err := http.NewRequest("DELETE", "/tasks/"+task.ID, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTaskByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}

func TestHandleTaskByIDMethodNotAllowed(t *testing.T) {
	setup()

	task := Task{ID: "1", Title: "Test Task", Description: "Test", Priority: 1}
	tasks[task.ID] = task

	req, err := http.NewRequest("PATCH", "/tasks/"+task.ID, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTaskByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestHandleTasksListAll(t *testing.T) {
	setup()

	// Manually create tasks with specified IDs
	tasks["1"] = Task{ID: "1", Title: "Test Task 1", Description: "Description 1", Priority: 1}
	tasks["2"] = Task{ID: "2", Title: "Test Task 2", Description: "Description 2", Priority: 2}

	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleTasks)

	// Call the handler with our request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Expect the response body to contain a JSON object with task IDs as keys.
	expectedTasks := map[string]Task{
		"1": {ID: "1", Title: "Test Task 1", Description: "Description 1", Priority: 1},
		"2": {ID: "2", Title: "Test Task 2", Description: "Description 2", Priority: 2},
	}

	var actualTasks map[string]Task
	if err := json.Unmarshal(rr.Body.Bytes(), &actualTasks); err != nil {
		t.Fatalf("Couldn't parse response body: %v", err)
	}

	if !compareTasks(expectedTasks, actualTasks) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedTasks)
	}
}

// Helper function to compare task maps
func compareTasks(expected, actual map[string]Task) bool {
	if len(expected) != len(actual) {
		return false
	}

	for key, expectedTask := range expected {
		if actualTask, found := actual[key]; found {
			if expectedTask != actualTask {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
