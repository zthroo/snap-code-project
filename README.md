# snap-code-project
Little coding project.

helpful sqlite:
.header on
.mode column

task_table schema:
CREATE TABLE task_table (
task_id INTEGER PRIMARY KEY AUTOINCREMENT,
user text NOT NULL,
task text NOT NULL,
status text CHECK( status IN('complete','incomplete'))
);

active_task_table schema:
CREATE TABLE active_task_table (
action_id INTEGER PRIMARY KEY AUTOINCREMENT,
task_id INTEGER NOT NULL,
user text NOT NULL,
timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
user_active_tasks INTEGER NOT NULL
);

trigger for adding task
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
