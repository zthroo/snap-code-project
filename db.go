package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type task struct {
	taskID int64
	user   string
	task   string
	status string
}

//TODO connect to taskDB func
func openTaskDB() (*sql.DB, error) {
	taskDB, err := sql.Open("sqlite3", "c:/users/pgold/side-projects/snap-code-project/local.db")
	if err != nil {
		//taskDB.Close()
		return nil, err
	}

	return taskDB, err
}

//TODO get all tasks from user
func getTasks(user string, taskDB *sql.DB) ([]task, error) {
	var tasks []task

	const query = `SELECT * FROM task_table WHERE user = ?`
	rows, err := taskDB.Query(query, user)
	if err != nil {
		return tasks, err
	}
	defer rows.Close()
	var task task
	for rows.Next() {
		err := rows.Scan(&task.taskID, &task.user, &task.task, &task.status)
		if err != nil {
			return nil, err
		}
		log.Println(task)
		tasks = append(tasks, task)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tasks, err
}

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
