package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("tasks/:user", getTasksWeb)
	router.GET("tasksCount/:user", getTaskCounts)
	router.GET("burndown/:user", getBurndown)

	router.POST("addTask", postNewTask)

	router.DELETE("deleteTask/:id", deleteTaskWeb)

	router.PUT("markTaskComplete/:id", updateTaskComplete)
	router.PUT("markTaskIncomplete/:id", updateTaskIncomplete)

	router.Run("localhost:8080")
}
