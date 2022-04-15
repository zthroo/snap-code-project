package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//TODO connect to taskDB func
func openTaskDB() (*sql.DB, error) {
	taskDB, err := sql.Open("sqlite3", "c:/users/pgold/side-projects/snap-code-project/local.db")
	if err != nil {
		//taskDB.Close()
		return nil, err
	}
	// TODO remove this commented out code
	// results, err := taskDB.Query("SELECT COUNT(*) FROM test WHERE id = 1")
	// var count int
	// for results.Next() {
	// 	err := results.Scan(&count)
	// 	if err != nil {
	// 		log.Println("error in scan")
	// 	}
	// }
	// log.Println("count: ", count)

	return taskDB, err
}

//TODO connect to auditDB func

//TODO get all tasks from user

//add a new task for a user
func addTask(user, task, status string, taskDB *sql.DB) int64 {
	const insertStmt = `INSERT INTO task_table (
		user,
		task,
		status
		) VALUES (?,?,?)`
	_, err := taskDB.Exec(insertStmt, user, task, status)
	if err != nil {
		panic(err) //TODO this is probably not how we want to handle this error since I think it will stop the service.
	}

	const getIDStmt = `select MAX(task_id) FROM task_table`
	var id int64
	err = taskDB.QueryRow(getIDStmt).Scan(&id)
	if err != nil {
		panic(err) //TODO this is probably not how we want to handle this error since I think it will stop the service.
	}
	return id
}

//TODO delete a task for a user

//TODO update a task to complete

//TODO update a task to incomplete

//TODO get # of complete and incomplete tasks for a user

//TODO get users count of active tasks with times
