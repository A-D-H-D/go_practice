package main

import "fmt"

func main() {
	// ASCII art border around welcome message
	fmt.Println("*******************************")
	fmt.Println("*  Welcome to our TodoList App  *")
	fmt.Println("*******************************")

	var shortGoLang = "Watch Go crash course"
	var fullGolang = "Watch Nana's Golang Full Course"
	var rewardDessert = "Reward myself with a donut"

	var taskItems = []string{shortGoLang, fullGolang, rewardDessert}
	printTasks(taskItems)

	// fmt.Println(taskItems[1:])
}

func printTasks(taskItems []string) {
	fmt.Println("\nList of my Todos:")

	for index, task := range taskItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}

}
