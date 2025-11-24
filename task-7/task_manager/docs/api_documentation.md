# Task Manager API with Authentication and Authorization

A RESTful API for managing tasks with JWT-based authentication and role-based authorization, built with Go and the Gin framework.

This project provides CRUD (Create, Read, Update, Delete) operations for a task management system with secure user authentication and authorization mechanisms.

**Note:** This API uses an in-memory data store. All tasks and users will be reset every time the application is restarted.

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or newer)

## Getting Started

1. **Navigate to the API directory:**
   ```sh
   cd task-7/task_manager
   ```

2. **Install dependencies:**
   ```sh
   go mod download
   ```

3. **Run the server:**
   ```sh
   go run main.go
   ```
   The server will start and listen on `http://localhost:8080`.

## Authentication

This API uses JSON Web Tokens (JWT) for authentication. To access protected endpoints, you must:

1. Register a new user account using `/auth/register`
2. Login using `/auth/login` to receive a JWT token
3. Include the token in the `Authorization` header for protected routes:
   ```
   Authorization: Bearer <your-token>
   ```

### User Roles

- **Admin**: Can create, update, and delete tasks. Can also promote other users to admin.
- **User**: Can only retrieve tasks (GET endpoints).

**Note:** The first user registered in the system automatically becomes an admin.

## API Endpoints

### Authentication Endpoints

#### Register a New User

- **URL:** `/auth/register`
- **Method:** `POST`
- **Description:** Creates a new user account. The first user registered becomes an admin automatically.
- **Headers:** `Content-Type: application/json`
- **Body (raw JSON):**
  ```json
  {
      "username": "john_doe",
      "password": "securepassword123"
  }
  ```
- **Response (201 Created):**
  ```json
  {
      "message": "User registered successfully",
      "user": {
          "id": 1,
          "username": "john_doe",
          "role": "admin"
      }
  }
  ```
- **Example `curl`:**
  ```sh
  curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username": "john_doe", "password": "securepassword123"}'
  ```

#### Login

- **URL:** `/auth/login`
- **Method:** `POST`
- **Description:** Authenticates a user and returns a JWT token.
- **Headers:** `Content-Type: application/json`
- **Body (raw JSON):**
  ```json
  {
      "username": "john_doe",
      "password": "securepassword123"
  }
  ```
- **Response (200 OK):**
  ```json
  {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "user": {
          "id": 1,
          "username": "john_doe",
          "role": "admin"
      }
  }
  ```
- **Example `curl`:**
  ```sh
  curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "john_doe", "password": "securepassword123"}'
  ```

### Task Endpoints

All task endpoints require authentication. Include the JWT token in the Authorization header.

#### Get All Tasks

- **URL:** `/tasks`
- **Method:** `GET`
- **Description:** Retrieves a list of all tasks. Accessible to all authenticated users.
- **Headers:** `Authorization: Bearer <token>`
- **Response (200 OK):**
  ```json
  {
      "tasks": [
          {
              "id": "1",
              "title": "Task 1",
              "description": "First task",
              "duedate": "2025-01-15T10:00:00Z",
              "status": "Pending"
          }
      ]
  }
  ```
- **Example `curl`:**
  ```sh
  curl http://localhost:8080/tasks \
  -H "Authorization: Bearer <your-token>"
  ```

#### Get Task by ID

- **URL:** `/tasks/:id`
- **Method:** `GET`
- **Description:** Retrieves a single task by its ID. Accessible to all authenticated users.
- **Headers:** `Authorization: Bearer <token>`
- **Response (200 OK):**
  ```json
  {
      "id": "1",
      "title": "Task 1",
      "description": "First task",
      "duedate": "2025-01-15T10:00:00Z",
      "status": "Pending"
  }
  ```
- **Example `curl`:**
  ```sh
  curl http://localhost:8080/tasks/1 \
  -H "Authorization: Bearer <your-token>"
  ```

#### Add a New Task (Admin Only)

- **URL:** `/tasks`
- **Method:** `POST`
- **Description:** Creates a new task. Only admins can create tasks.
- **Headers:** 
  - `Content-Type: application/json`
  - `Authorization: Bearer <token>`
- **Body (raw JSON):**
  ```json
  {
      "id": "4",
      "title": "New Task from API",
      "description": "A task created via POST request.",
      "duedate": "2025-12-01T15:00:00Z",
      "status": "Pending"
  }
  ```
- **Response (201 Created):**
  ```json
  {
      "message": "Task Created"
  }
  ```
- **Example `curl`:**
  ```sh
  curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{"id": "4", "title": "New Task", "description": "A new task", "duedate": "2025-12-01T15:00:00Z", "status": "Pending"}'
  ```

#### Update a Task (Admin Only)

- **URL:** `/tasks/:id`
- **Method:** `PUT`
- **Description:** Updates an existing task. Only admins can update tasks.
- **Headers:** 
  - `Content-Type: application/json`
  - `Authorization: Bearer <token>`
- **Body (raw JSON):**
  ```json
  {
      "title": "Updated Task Title",
      "description": "This task has been updated.",
      "status": "In Progress"
  }
  ```
- **Response (200 OK):**
  ```json
  {
      "message": "Task Updated"
  }
  ```
- **Example `curl`:**
  ```sh
  curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{"title": "Updated Title"}'
  ```

#### Delete a Task (Admin Only)

- **URL:** `/tasks/:id`
- **Method:** `DELETE`
- **Description:** Deletes a task by its ID. Only admins can delete tasks.
- **Headers:** `Authorization: Bearer <token>`
- **Response (200 OK):**
  ```json
  {
      "message": "Task removed"
  }
  ```
- **Example `curl`:**
  ```sh
  curl -X DELETE http://localhost:8080/tasks/1 \
  -H "Authorization: Bearer <your-token>"
  ```

### Admin Endpoints

#### Promote User to Admin

- **URL:** `/admin/promote`
- **Method:** `POST`
- **Description:** Promotes a user to admin role. Only admins can promote users.
- **Headers:** 
  - `Content-Type: application/json`
  - `Authorization: Bearer <token>`
- **Body (raw JSON):**
  ```json
  {
      "username": "jane_doe"
  }
  ```
- **Response (200 OK):**
  ```json
  {
      "message": "User promoted to admin successfully"
  }
  ```
- **Example `curl`:**
  ```sh
  curl -X POST http://localhost:8080/admin/promote \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{"username": "jane_doe"}'
  ```

## Error Responses

### 400 Bad Request
```json
{
    "error": "Invalid request format"
}
```

### 401 Unauthorized
```json
{
    "error": "Authorization header required"
}
```
or
```json
{
    "error": "Invalid or expired token"
}
```

### 403 Forbidden
```json
{
    "error": "Admin access required"
}
```

### 404 Not Found
```json
{
    "error": "Task not found"
}
```

## Security Features

1. **Password Hashing**: User passwords are hashed using bcrypt before storage.
2. **JWT Tokens**: Secure token-based authentication for API access.
3. **Role-Based Access Control**: Different permissions for admin and regular users.
4. **Protected Routes**: All task endpoints require valid authentication.
5. **Admin-Only Operations**: Create, update, and delete operations are restricted to admins.

## Testing the API

### Complete Workflow Example

1. **Register the first user (becomes admin):**
   ```sh
   curl -X POST http://localhost:8080/auth/register \
   -H "Content-Type: application/json" \
   -d '{"username": "admin", "password": "admin123"}'
   ```

2. **Login to get token:**
   ```sh
   curl -X POST http://localhost:8080/auth/login \
   -H "Content-Type: application/json" \
   -d '{"username": "admin", "password": "admin123"}'
   ```
   Save the token from the response.

3. **Create a task (admin only):**
   ```sh
   curl -X POST http://localhost:8080/tasks \
   -H "Content-Type: application/json" \
   -H "Authorization: Bearer <your-token>" \
   -d '{"id": "1", "title": "My Task", "description": "Task description", "duedate": "2025-12-31T23:59:59Z", "status": "Pending"}'
   ```

4. **Get all tasks (any authenticated user):**
   ```sh
   curl http://localhost:8080/tasks \
   -H "Authorization: Bearer <your-token>"
   ```

5. **Register a regular user:**
   ```sh
   curl -X POST http://localhost:8080/auth/register \
   -H "Content-Type: application/json" \
   -d '{"username": "user1", "password": "user123"}'
   ```

6. **Promote user to admin:**
   ```sh
   curl -X POST http://localhost:8080/admin/promote \
   -H "Content-Type: application/json" \
   -H "Authorization: Bearer <admin-token>" \
   -d '{"username": "user1"}'
   ```

## Notes

- The JWT secret key is hardcoded in the middleware. In production, use environment variables.
- All data is stored in memory and will be lost when the server restarts.
- Token expiration is not currently implemented but can be added by uncommenting the expiration logic in the middleware.

