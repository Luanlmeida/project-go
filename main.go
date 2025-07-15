package main

import "fmt"

func main() {
	var fullGolang = "Watch Golang Full course"
	var shortGolang = "Watch Go crash course"
	var rewardDessert = "Reward my self with a donut"

	taskItens := []string{fullGolang, shortGolang, rewardDessert}

	fmt.Println("##### Welcome to our Todolist App! #####")

	fmt.Println("List of my Todos")

	for index, task := range taskItens {
		fmt.Println(index+1, ".", task)
		fmt.Printf("%d: %s\n", index+1, task)
	}
}
