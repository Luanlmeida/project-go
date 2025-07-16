package main

import "fmt"

func main() {

	fmt.Println("##### Welcome to our Todolist App! #####")

	var fullGolang = "Watch Golang Full course"
	var shortGolang = "Watch Go crash course"
	var rewardDessert = "Reward my self with a donut"
	var taskItems = []string{fullGolang, shortGolang, rewardDessert}

	printTasks(taskItems)
	fmt.Println()

	taskItems = addTaks(taskItems, "Go for a run")
	taskItems = addTaks(taskItems, "Practice coding in Go")

	fmt.Println()
	printTasks(taskItems)
}

func printTasks(taskItens []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItens {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTaks(taskItens []string, newTask string) []string {
	var updatedTaskItems = append(taskItens, newTask)
	return updatedTaskItems
}
