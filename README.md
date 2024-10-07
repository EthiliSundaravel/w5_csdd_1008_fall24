# Go Task CRUD API

A simple Task Management API built in Go, supporting CRUD operations for tasks, each identified by a unique ID and containing a title, description, and status ("pending" or "completed").


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
- Go version 1.23 installed.

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

- **Response (201 Created)**:
   
   ```json
   {
      "id": 1,
      "title": "Task Title",
      "description": "Task Description",
      "status": "pending"
   }

### 2. Get All Tasks (GET /tasks)

- **Method:** GET
- **Endpoint:** /tasks

- **Response** (200 OK):

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
- **Response** (200 OK):

   ```json
   {
      "id": 1,
      "title": "Task Title 1",
      "description": "Task Description 1",
      "status": "pending"
   }
   ```
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
 
- **Response (200 OK)**:

   ```json
   {
      "id": 1,
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "status": "completed"
   }
   ```
### 5. Delete a Tasks by ID (DELETE /tasks/{id})

- **Method:** DELETE
- **Endpoint:** /tasks/{id}

- **Response** (200 Created):

   ```json
   {
      "message": "Task deleted successfully"
   }

## Code Explanation

### 1. Task Struct
Defines the structure of a task:

```go
type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"` // "pending" or "completed"
}
```
### 2. In-Memory Storage

Tasks are stored in an in-memory slice, with an auto-incrementing nextID for unique task IDs:

```go
var (
    tasks    = []Task{}    // In-memory slice to store tasks
    nextID   = 1           // Auto-incrementing task ID
    tasksMux sync.Mutex     // Mutex for concurrent access
)
```
### 3. Status Validation

To ensure that the status field only accepts valid values ("pending" or "completed"), a ValidateStatus function is implemented. The validStatuses map contains the allowed statuses, and the function checks if the provided status exists in this map.

```go
func ValidateStatus(status string) bool {
    _, exists := validStatuses[status]
    return exists
}
```
### 4. Handlers

The API provides the following handlers for managing tasks:

- **POST `/tasks`**: Create a new task.
- **GET `/tasks`**: Retrieve all tasks.
- **GET `/tasks/{id}`**: Get a specific task by its ID.
- **PUT `/tasks/{id}`**: Update an existing task.
- **DELETE `/tasks/{id}`**: Delete a task by its ID.

### 5. HTTP Server

Sets up the server and routes:

```go
http.HandleFunc("/tasks", ...)
http.HandleFunc("/tasks/{id}", ...)
log.Fatal(http.ListenAndServe(":8080", nil))
```

## Conclusion
This API serves as a basic framework for task management applications, utilizing Go's net/http package for building RESTful services.