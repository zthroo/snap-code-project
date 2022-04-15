package main

import (
	"log"
)

func main() {

	// hello world web server
	con, err := openTaskDB()
	if err != nil {
		log.Println("Error connecting to task DB.")
	}

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

	completeCount, incompleteCount, err := getCompleteAndIncompleteCount("person@email.com", con)
	if err != nil {
		log.Println("error: ", err)
	} else {
		log.Println("complete count:", completeCount)
		log.Println("incomplete count:", incompleteCount)
	}

	timeCounts, err := getTimeCounts("person@email.com", con)
	if err != nil {
		log.Println("error:", err)
	} else {
		log.Println("timeCounts:", timeCounts)
	}

	tasks, err := getTasks("person@email.com", con)
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println("Tasks: ", tasks)

	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// }

	// http.HandleFunc("/hello", helloHandler)
	// log.Println("Listting for requests at http://localhost:8000/hello")
	// log.Fatal(http.ListenAndServe(":8000", nil))
}
