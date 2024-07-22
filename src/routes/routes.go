package routes

import (
    "github.com/gin-gonic/gin"
    "todo-app/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/tasks", controllers.GetTasks)
    r.POST("/tasks", controllers.CreateTask)
    r.PUT("/tasks/:id", controllers.UpdateTask)
    r.DELETE("/tasks/:id", controllers.DeleteTask)
    r.GET("/tasks/:id", controllers.ShowTask)
    r.GET("/tasks/status", controllers.GetTasksByStatus)

    return r
}
