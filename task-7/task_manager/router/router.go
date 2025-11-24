package router

import (
	"a2sv-backend/task_manager/controllers"
	"a2sv-backend/task_manager/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Public authentication routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	// Protected routes (require authentication)
	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		// Task routes - GET endpoints accessible to all authenticated users
		taskRoutes := protectedRoutes.Group("/tasks")
		{
			taskRoutes.GET("", controllers.GetTasks)
			taskRoutes.GET("/:id", controllers.GetTask)

			// Admin-only task routes
			adminTaskRoutes := taskRoutes.Group("")
			adminTaskRoutes.Use(middleware.AdminOnlyMiddleware())
			{
				adminTaskRoutes.POST("", controllers.AddTask)
				adminTaskRoutes.PUT("/:id", controllers.UpdateTask)
				adminTaskRoutes.DELETE("/:id", controllers.DeleteTask)
			}
		}

		// Admin-only routes
		adminRoutes := protectedRoutes.Group("/admin")
		adminRoutes.Use(middleware.AdminOnlyMiddleware())
		{
			adminRoutes.POST("/promote", controllers.Promote)
		}
	}

	return router
}

