package main

import "fmt"

func main() {

	fmt.Println("##### Welcome to our Todolist App! #####")

	var fullGolang = "Watch Golang Full course"
	var shortGolang = "Watch Go crash course"
	var rewardDessert = "Reward my self with a donut"
	var taskItens = []string{fullGolang, shortGolang, rewardDessert}

	printTasks(taskItens)
	fmt.Println()
	addTaks(taskItens, "Go for a run")
}

func printTasks(taskItens []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItens {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTaks(taskItens []string, newTask string) {
	var updatedTaskItems = append(taskItens, newTask)
	printTasks(updatedTaskItems)
}
