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
