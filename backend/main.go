// Package main contains the backend server logic for TaskBoard, including
// database initialization, environment variable handling, and application startup.
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	initDB()

	r := gin.Default()

	// Simple CORS for local dev
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", getEnv("FRONTEND_ORIGIN", "http://localhost:5173"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		api.GET("/tasks", getTasks)
		api.POST("/tasks", createTask)
		api.PUT("/tasks/:id", updateTask)
		api.DELETE("/tasks/:id", deleteTask)
	}

	port := getEnv("PORT", "8080")
	log.Printf("🚀 Backend running on :%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
