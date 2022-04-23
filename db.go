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

type timeCount struct {
	time      string
	taskCount int
}

// connect to taskDB func
func openTaskDB() (*sql.DB, error) {
	taskDB, err := sql.Open("sqlite3", "c:/users/pgold/side-projects/snap-code-project/local.db")
	if err != nil {
		//taskDB.Close()
		return nil, err
	}

	return taskDB, err
}

// get all tasks from user
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
		log.Println(task) //TODO comment out for final
		tasks = append(tasks, task)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tasks, err
}

// add a new task for a user
func addTask(user, task string, taskDB *sql.DB) (int64, error) {
	const insertStmt = `INSERT INTO task_table (
		user,
		task,
		status
		) VALUES (?,?,'incomplete')`
	_, err := taskDB.Exec(insertStmt, user, task)
	if err != nil {
		return 0, err
	}

	const getIDStmt = `select MAX(task_id) FROM task_table`
	var id int64
	err = taskDB.QueryRow(getIDStmt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}

// delete specific task
func deleteTask(id int64, taskDB *sql.DB) error {
	const deleteStmt = `DELETE FROM task_table WHERE task_id = ?`
	_, err := taskDB.Exec(deleteStmt, id)
	return err
}

// update a task to complete
func markTaskComplete(id int64, taskDB *sql.DB) error {
	const updateStmt = `UPDATE task_table SET status = 'complete' WHERE task_id = ? AND status = 'incomplete'`
	_, err := taskDB.Exec(updateStmt, id)
	return err
}

// update a task to incomplete
func markTaskIncomplete(id int64, taskDB *sql.DB) error {
	const updateStmt = `UPDATE task_table SET status = 'incomplete' WHERE task_id = ? AND status = 'complete'`
	_, err := taskDB.Exec(updateStmt, id)
	return err
}

// get # of complete and incomplete tasks for a user
func getCompleteAndIncompleteCount(user string, taskDB *sql.DB) (int, int, error) {
	const completeQuery = `SELECT COUNT(*) FROM task_table WHERE user = ? AND status = 'complete'`
	const incompleteQuery = `SELECT COUNT(*) FROM task_table WHERE user = ? AND status = 'incomplete'`

	var completeCount int
	err := taskDB.QueryRow(completeQuery, user).Scan(&completeCount)
	if err != nil {
		return 0, 0, err
	}
	var incompleteCount int
	err = taskDB.QueryRow(incompleteQuery, user).Scan(&incompleteCount)
	if err != nil {
		return 0, 0, err
	}

	return completeCount, incompleteCount, err
}

// get users count of active tasks with times
func getTimeCounts(user string, taskDB *sql.DB) ([]timeCount, error) {
	var timeCounts []timeCount
	const query = `SELECT 
			timestamp, 
			user_active_tasks 
		FROM 
			active_task_table 
		WHERE 
			user = ? 
		ORDER BY 
			timestamp ASC`

	rows, err := taskDB.Query(query, user)
	if err != nil {
		return timeCounts, err
	}
	defer rows.Close()
	var timeCount timeCount
	for rows.Next() {
		err := rows.Scan(&timeCount.time, &timeCount.taskCount)
		if err != nil {
			return nil, err
		}
		log.Println(timeCount) //TODO comment out for final
		timeCounts = append(timeCounts, timeCount)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return timeCounts, err
}
