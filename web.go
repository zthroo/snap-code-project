package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getTasksWeb returns all the current tasks of the requested user as a json array
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

// getTaskCount returns the count of the requested user's complete and incomplete tasks as a json
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

// getBurndown returns an array of the requested user's active tasks counts with timestamps as a json array
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

// postNewTask expects a json containing the user and the task, checks that it contains both of those, inserts the task into the database, and returns the task id assigned.
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

// deleteTaskWeb deletes the requested task from the database and returns the id of the deleted task
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

// updateTaskComplete updates the status of the task with id matching the given id to complete
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

// updateTaskIncomplete updates the status of the task with id matching the given id to incomplete
func updateTaskIncomplete(c *gin.Context) {
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
	err = markTaskIncomplete(id, taskDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, id)
}
