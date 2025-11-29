package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getTasks(c *gin.Context) {
	var tasks []Task
	if err := DB.Order("created_at desc").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

type CreateTaskInput struct {
	Title string `json:"title" binding:"required,min=1,max=200"`
}

func createTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	task := Task{
		Title:     input.Title,
		Completed: false,
	}
	if err := DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

type UpdateTaskInput struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

func updateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var task Task
	if err := DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Completed != nil {
		task.Completed = *input.Completed
	}

	if err := DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := DB.Delete(&Task{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
		return
	}

	c.Status(http.StatusNoContent)
}
