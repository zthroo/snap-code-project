package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to db
	// con, err := openTaskDB()
	// if err != nil {
	// 	log.Println("Error connecting to task DB.")
	// }

	router := gin.Default()
	router.GET("tasks/:user", getTasksWeb)
	router.GET("tasksCount/:user", getTaskCounts)
	router.GET("burndown/:user", getBurndown)

	router.POST("addTask", postNewTask)

	router.DELETE("deleteTask/:id", deleteTaskWeb)

	router.PUT("markTaskComplete/:id", updateTaskComplete)

	router.Run("localhost:8080")

	// id, err := addTask("person@email.com", "sample task 5", con)
	// if err != nil {
	// 	log.Println("error: ", err)
	// }
	// log.Println("ID: ", id)

	// err = deleteTask(6, con)
	// if err != nil {
	// 	log.Println("error: ", err)
	// }

	// err = markTaskComplete(9, con)
	// if err != nil {
	// 	log.Println("error: ", err)
	// }

	// completeCount, incompleteCount, err := getCompleteAndIncompleteCount("person@email.com", con)
	// if err != nil {
	// 	log.Println("error: ", err)
	// } else {
	// 	log.Println("complete count:", completeCount)
	// 	log.Println("incomplete count:", incompleteCount)
	// }

	// timeCounts, err := getTimeCounts("person@email.com", con)
	// if err != nil {
	// 	log.Println("error:", err)
	// } else {
	// 	log.Println("timeCounts:", timeCounts)
	// }

	// tasks, err := getTasks("person@email.com", con)
	// if err != nil {
	// 	log.Println("error: ", err)
	// }
	// log.Println("Tasks: ", tasks)

	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// }

	// http.HandleFunc("/hello", helloHandler)
	// log.Println("Listting for requests at http://localhost:8000/hello")
	// log.Fatal(http.ListenAndServe(":8000", nil))
}
