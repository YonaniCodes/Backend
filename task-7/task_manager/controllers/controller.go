package controllers

import (
	"a2sv-backend/task_manager/data"
	"a2sv-backend/task_manager/middleware"
	"a2sv-backend/task_manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User Authentication Controllers

// Register handles user registration
func Register(ctx *gin.Context) {
	var req models.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := data.RegisterUser(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Don't send password in response
	user.Password = ""
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    user,
	})
}

// Login handles user authentication and returns JWT token
func Login(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := data.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, models.LoginResponse{
		Token: token,
		User:  *user,
	})
}

// Promote handles admin promotion of users
func Promote(ctx *gin.Context) {
	var req models.PromoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.PromoteUser(req.Username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User promoted to admin successfully"})
}

// Task Controllers

// GetTasks retrieves all tasks (accessible by authenticated users)
func GetTasks(ctx *gin.Context) {
	tasks := data.GetAllTasks()
	ctx.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

// GetTask retrieves a task by ID (accessible by authenticated users)
func GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := data.GetTaskByID(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, task)
}

// AddTask creates a new task (admin only)
func AddTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.AddTask(newTask)
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Task Created"})
}

// UpdateTask updates an existing task (admin only)
func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := data.UpdateTask(id, updatedTask); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task Updated"})
}

// DeleteTask deletes a task (admin only)
func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := data.DeleteTask(id); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Task removed"})
}

