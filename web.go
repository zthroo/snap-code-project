package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTasksWeb(c *gin.Context) {
	user := c.Param("user")
	taskDB, err := openTaskDB()
	tasks, err := getTasks(user, taskDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskCounts(c *gin.Context) {
	user := c.Param("user")
	taskDB, err := openTaskDB()
	tasksCount, err := getCompleteAndIncompleteCount(user, taskDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, tasksCount)
}
