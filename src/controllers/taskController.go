package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "todo-app/models"
)

func GetTasks(c *gin.Context) {
    var tasks []models.Task
    models.DB.Find(&tasks)
    c.JSON(http.StatusOK, gin.H{"task": tasks})
}

func CreateTask(c *gin.Context) {
    var newTask models.InputTask
    if err := c.BindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.DB.Create(&newTask)
    c.JSON(http.StatusCreated, gin.H{"message": "Task created", "task": newTask})
}

func UpdateTask(c *gin.Context) {
    var task models.InputTask
    if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    if err := c.BindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.DB.Save(&task)
    c.JSON(http.StatusOK, gin.H{"message": "Task updated", "task": task})
}

func DeleteTask(c *gin.Context) {
    var task models.Task
    if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    models.DB.Delete(&task)
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func ShowTask(c *gin.Context) {
    var task models.Task
    if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"task": task})
}

//特定のstatusのTask一覧を取得。getTasksの中にまとめることもできそう。
func GetTasksByStatus(c *gin.Context) {
    var tasks []models.Task
    status := c.Query("status")

    if err := models.DB.Where("status = ?", status).Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"task": tasks})
}