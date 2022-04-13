package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//TODO connect to taskDB func
func openTaskDB() (*sql.DB, error) {
	taskDB, err := sql.Open("sqlite3", "c:/users/pgold/side-projects/snap-code-project/local.db")
	if err != nil {
		//taskDB.Close()
		return nil, err
	}
	results, err := taskDB.Query("SELECT COUNT(*) FROM test WHERE id = 1")
	var count int
	for results.Next() {
		err := results.Scan(&count)
		if err != nil {
			log.Println("error in scan")
		}
	}
	log.Println("count: ", count)

	return taskDB, err
}

//TODO connect to auditDB func

//TODO get all tasks from user

//TODO add a new task for a user
func addTask(taskId, user, task, status string, taskDB *sql.DB) {
	const insertStmt = `INSERT INTO task_table (
		task_id,
		user,
		task_string,
		status
		) VALUES (?,?,?,?)`
	_, err := taskDB.Exec(insertStmt, taskId, user, task, status)
	if err != nil {
		panic(err) //TODO this is probably not how we want to handle this error since I think it will stop the service.
	}
}

//TODO delete a task for a user

//TODO update a task to complete

//TODO update a task to incomplete

//TODO get # of complete and incomplete tasks for a user

//TODO get users count of active tasks with times
