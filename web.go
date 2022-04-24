package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getTasksWeb(c *gin.Context) {
	user := c.Param("user")
	taskDB, err := openTaskDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	tasks, err := getTasks(user, taskDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func getTaskCounts(c *gin.Context) {
	user := c.Param("user")
	taskDB, err := openTaskDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	tasksCount, err := getCompleteAndIncompleteCount(user, taskDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, tasksCount)
}

func getBurndown(c *gin.Context) {
	user := c.Param("user")
	taskDB, err := openTaskDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	burndown, err := getTimeCounts(user, taskDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
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

func deleteTaskWeb(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	taskDB, err := openTaskDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = deleteTask(id, taskDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, id)
}

func updateTaskComplete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 0, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	taskDB, err := openTaskDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = markTaskComplete(id, taskDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, id)
}
