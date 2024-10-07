package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Task struct represents a basic task entity
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // "pending" or "completed"
}

var (
	tasks     = []Task{}       // In-memory slice to store tasks
	nextID    = 1              // Auto-incrementing task ID
	tasksMux  sync.Mutex       // Mutex for concurrent access
)

// Valid statuses for a task
var validStatuses = map[string]bool{
	"pending":   true,
	"completed": true,
}

// ValidateStatus checks if the status is valid
func ValidateStatus(status string) bool {
	_, exists := validStatuses[status]
	return exists
}
// Handler for creating a new task (POST /tasks)
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Validate status
	if !ValidateStatus(newTask.Status) {
		http.Error(w, "Invalid status field", http.StatusBadRequest)
		return
	}

	tasksMux.Lock()
	defer tasksMux.Unlock()

	newTask.ID = nextID
	nextID++
	tasks = append(tasks, newTask)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// Handler for getting all tasks (GET /tasks)
func getAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	tasksMux.Lock()
	defer tasksMux.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Handler for getting a task by ID (GET /tasks/{id})
func getTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/tasks/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	tasksMux.Lock()
	defer tasksMux.Unlock()

	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

// Handler for updating a task (PUT /tasks/{id})
func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/tasks/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Validate status
	if !ValidateStatus(updatedTask.Status) {
		http.Error(w, "Invalid status field", http.StatusBadRequest)
		return
	}
	
	tasksMux.Lock()
	defer tasksMux.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id // Keep the original ID
			tasks[i] = updatedTask
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

// Handler for deleting a task (DELETE /tasks/{id})
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/tasks/"):]

	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	tasksMux.Lock()
	defer tasksMux.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			// Delete the task by removing it from the slice
			tasks = append(tasks[:i], tasks[i+1:]...) // Remove task
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}


// Main function to set up the routes and start the HTTP server
func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllTasksHandler(w, r)
		case http.MethodPost:
			createTaskHandler(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTaskByIDHandler(w, r)
		case http.MethodPut:
			updateTaskHandler(w, r)
		case http.MethodDelete:
			deleteTaskHandler(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
