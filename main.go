package main

import (
	"fmt"
	"net/http"
)

var fullGolang = "Watch Golang Full course"
var shortGolang = "Watch Go crash course"
var rewardDessert = "Reward my self with a donut"
var taskItems = []string{fullGolang, shortGolang, rewardDessert}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}

func helloUser(writter http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our Todolist App!"
	fmt.Fprintln(writter, greeting)
}

func showTasks(writter http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writter, task)
	}
}
