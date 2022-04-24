package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTasksWeb(c *gin.Context) {
	user := c.Param("user")
	taskDB, err := openTaskDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	tasks, err := getTasks(user, taskDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskCounts(c *gin.Context) {
	user := c.Param("user")
	taskDB, err := openTaskDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	tasksCount, err := getCompleteAndIncompleteCount(user, taskDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, tasksCount)
}

func getBurndown(c *gin.Context) {
	user := c.Param("user")
	taskDB, err := openTaskDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	burndown, err := getTimeCounts(user, taskDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, burndown)
}

type CreateTaskInput struct {
	User string `json:"user" binding:"required"`
	Task string `json:"task" binding:"required"`
}

func postNewTask(c *gin.Context) {
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	taskDB, err := openTaskDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	id, err := addTask(input.User, input.Task, taskDB)
	c.IndentedJSON(http.StatusOK, gin.H{"taskId": id})
}
