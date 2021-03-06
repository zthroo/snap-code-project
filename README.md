# snap-code-project
The backend REST API and db implementation for a simple ToDo app.

## GET
* GET get list of existing to-do tasks - endpoint: `/tasks/<user>`
* GET users counts with times - endpoint: `/burndown/<user>`
* GET # of complete and incomplete teaks for a user - endpoint: `/tasksCount/<user>`

## POST
* POST add a new to-do task - endpoint: `/addTask` body should be `{"user":<user>,"task":<task>}` 
    * example curl: `curl -X POST localhost:8080/addTask -H "Content-Type: application/json" -d "{\"user\":\"specialperson@email.com\",\"task\":\"a new task\"}"`

## PUT
* PUT mark task as complete - endpoint: `markTaskComplete/<taskID>`
    * example curl: `curl -X "PUT" localhost:8080/markTaskComplete/11`
* PUT mark task as incomplete endpoint: `markTaskIncomplete/<taskID>`
    * example curl: `curl -X "PUT" localhost:8080/markTaskIncomplete/11`

## DELETE
* DELETE a task - endpoint: `/deleteTask/<taskID>`
    * example curl: `curl -X "DELETE" localhost:8080/deleteTask/15`


helpful sqlite:
* .header on
* .mode column

## task_table schema:
```
CREATE TABLE task_table (
task_id INTEGER PRIMARY KEY AUTOINCREMENT,
user text NOT NULL,
task text NOT NULL,
status text CHECK( status IN('complete','incomplete'))
);
```

## active_task_table schema:
```
CREATE TABLE active_task_table (
action_id INTEGER PRIMARY KEY AUTOINCREMENT,
task_id INTEGER NOT NULL,
user text NOT NULL,
timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
user_active_tasks INTEGER NOT NULL
);
```

## trigger for adding task
```
CREATE TRIGGER Added_Task
    AFTER INSERT ON task_table
BEGIN
        INSERT INTO active_task_table (
            task_id,
            user,
            user_active_tasks
        )
        VALUES
        (
            NEW.task_id,
            NEW.user,
            (CASE WHEN ((SELECT user 
                        FROM active_task_table
                        WHERE user = NEW.user) IS NOT NULL)
            THEN ((SELECT user_active_tasks
                    FROM active_task_table
                    WHERE user = NEW.user
                    ORDER BY timestamp DESC
                    LIMIT 1) + 1)
            ELSE 1
            END)
        );
END;
```

## trigger for delete task
```
CREATE TRIGGER Deleted_Task
    AFTER DELETE ON task_table
    WHEN OLD.status = 'incomplete'
BEGIN
        INSERT INTO active_task_table (
            task_id,
            user,
            user_active_tasks
        )
        VALUES
        (
            OLD.task_id,
            OLD.user,
            ((SELECT user_active_tasks
                FROM active_task_table
                WHERE user = OLD.user
                ORDER BY timestamp DESC
                LIMIT 1) - 1)
        );
END;
```

## trigger for update task to complete
```
CREATE TRIGGER Completed_Task
    AFTER UPDATE on task_table
    WHEN new.status = 'complete'
BEGIN
        INSERT INTO active_task_table (
            task_id,
            user,
            user_active_tasks
        )
        VALUES
        (
            OLD.task_id,
            OLD.user,
            ((SELECT user_active_tasks
                FROM active_task_table
                WHERE user = OLD.user
                ORDER BY timestamp DESC
                LIMIT 1) - 1)
        );
END;
```

## trigger for update task to incomplete
```
CREATE TRIGGER Incompleted_Task
    AFTER UPDATE on task_table
    WHEN new.status = 'incomplete'
BEGIN
        INSERT INTO active_task_table (
            task_id,
            user,
            user_active_tasks
        )
        VALUES
        (
            OLD.task_id,
            OLD.user,
            ((SELECT user_active_tasks
                FROM active_task_table
                WHERE user = OLD.user
                ORDER BY timestamp DESC
                LIMIT 1) + 1)
        );
END;
```