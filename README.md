# Go Task CRUD API

This is a simple RESTful CRUD (Create, Read, Update, Delete) API built using Go version 1.23 without any third-party packages. It uses Go's built-in HTTP server, JSON encoding/decoding, and concurrency mechanisms like mutexes for handling in-memory task storage.

## Table of Contents
- [Features](#features)
- [Project Structure](#project-structure)
- [Routes](#routes)
- [Installation and Running](#installation-and-running)
- [API Usage](#api-usage)
- [Code Explanation](#code-explanation)

## Features
This API provides basic CRUD operations for managing tasks:
1. **Create** a new task.
2. **Read** a single task by ID or list all tasks.
3. **Update** an existing task by ID.
4. **Delete** a task by ID.

## Project Structure

```
W5_CSDD_1008_FALL24/
|-- main.go
```

## Routes
The API exposes the following endpoints:
- `GET /tasks` - Retrieve all tasks.
- `POST /tasks` - Create a new task.
- `GET /tasks/{id}` - Retrieve a task by ID.
- `PUT /tasks/{id}` - Update a task by ID.
- `DELETE /tasks/{id}` - Delete a task by ID.

## Installation and Running

### Prerequisites:
- Go version 1.23 installed on your machine.

### Running the server:
1. Clone the repository or copy the `main.go` file.
2. Open a terminal and navigate to the project directory.
3. Run the following command to start the server:
   ```bash
   go run main.go

The server will start on http://localhost:8080.

## API Usage

### 1. Create a Task (POST /tasks)

- **Method:** POST
- **Endpoint:** /tasks
- **Request Body** (JSON):
  
  ```json
  {
      "title": "Task Title",
      "description": "Task Description",
      "status": "pending"
  }
  ```

- **Response**:

- **201 Created**
   
   ```json
   {
      "id": 1,
      "title": "Task Title",
      "description": "Task Description",
      "status": "pending"
   }
   ```
- **400 Bad Request** (if the status is invalid)

### 2. Get All Tasks (GET /tasks)

- **Method:** GET
- **Endpoint:** /tasks

- **Response** (201 Created):

```json
[
            {
                "id": 1,
                "title": "Task Title 1",
                "description": "Task Description 1",
                "status": "pending"
            },
            {
                "id": 2,
                "title": "Task Title 2",
                "description": "Task Description 2",
                "status": "completed"
            }
        ]
```

### 3. Get a Tasks by ID (GET /tasks/{id})

- **Method:** GET
- **Endpoint:** /tasks/{id}
- **Response** (201 Created):

- **200 OK**
   ```json
   {
      "id": 1,
      "title": "Task Title 1",
      "description": "Task Description 1",
      "status": "pending"
   }
   ```
- **404 Not Found** (if the task with the given ID does not exist)

### 4. Update a Tasks by ID (PUT /tasks/{id})

- **Method:** PUT
- **Endpoint:** /tasks/{id}
- **Request Body** (JSON):
  
  ```json
  {
        "title": "Updated Task Title",
        "description": "Updated Task Description",
        "status": "completed"
  }
 
- **Response**:

- **200 OK**
   ```json
   {
      "id": 1,
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "status": "completed"
   }
   ```
- **400 Bad Request** (if the status is invalid)
- **404 Not Found** (if the task with the given ID does not exist)

### 5. Delete a Tasks by ID (DELETE /tasks/{id})

- **Method:** DELETE
- **Endpoint:** /tasks/{id}

- **Response** (201 Created):

- **200 OK**
   ```json
   {
      "message": "Task deleted successfully"
   }

- **404 Not Found** (if the task with the given ID does not exist)

## Code Explanation

### 1. Task Struct
The `Task` struct represents a basic task entity in our application. It contains the following fields:

- **ID**: A unique identifier for each task.
- **Title**: The title of the task.
- **Description**: A detailed description of the task.
- **Status**: The current status of the task, which can be either "pending" or "completed".

```go
type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"` // "pending" or "completed"
}
```

### 2. In-Memory Storage

We use an in-memory slice (tasks) to store the tasks. The nextID variable is used to auto-increment task IDs, ensuring each task has a unique identifier.

```go
var (
    tasks    = []Task{}    // In-memory slice to store tasks
    nextID   = 1           // Auto-incrementing task ID
    tasksMux sync.Mutex     // Mutex for concurrent access
)
```
The tasksMux mutex is employed to ensure thread-safe access to the tasks slice, allowing multiple requests to modify the slice without causing data races.

### 3. Status Validation

To ensure that the status field only accepts valid values ("pending" or "completed"), a ValidateStatus function is implemented:

```go
func ValidateStatus(status string) bool {
    _, exists := validStatuses[status]
    return exists
}
```

The validStatuses map contains the allowed statuses, and the function checks if the provided status exists in this map.

### 4. Handlers

We define separate handlers for each CRUD operation related to tasks.

- **Create Task (POST /tasks)**: Handles task creation by decoding the request body, validating the status, and adding the task to the in-memory storage.

- **Get All Tasks (GET /tasks)**: Fetches and returns a list of all available tasks.

- **Get Task by ID (GET /tasks/{id})**: Retrieves a specific task by its ID.

- **Update Task (PUT /tasks/{id})**: Updates the details of an existing task, including validating the status field.

- **Delete Task (DELETE /tasks/{id})**: Removes a task by its ID from the in-memory storage.

### 5. HTTP Server

We use Go's built-in net/http package to set up an HTTP server and define routes.

- http.HandleFunc("/tasks", ...): This defines the route for listing and creating tasks.
- http.HandleFunc("/tasks/{id}", ...): This defines routes for getting, updating, and deleting tasks by ID.

The server is started using:
```go
log.Fatal(http.ListenAndServe(":8080", nil))
```
This binds the server to port 8080 and listens for incoming requests.

## Conclusion
This CRUD API is a simple demonstration of using Go 1.23's built-in packages to create a basic API for managing tasks. The API supports all CRUD operations and ensures thread safety with mutexes for concurrent access to in-memory data.