package main

import (
	"fmt"
	"net/http"
)

var fullGolang = "Watch Nana's Golang Full Course"
var shortGolang = "Watch Go crash course"
var rewardDessert = "Reward myself with a donut"
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {
	fmt.Println("###### Welcome to our Todolist App! ######")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)

	printTasks(taskItems)
	fmt.Println()

	taskItems = addTask(taskItems, "Go for a run")
	taskItems = addTask(taskItems, "Practice coding in Go")

	fmt.Println("Updated list")
	printTasks(taskItems)

}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	greeting := "Hello user. Welcome to our Todolist App!"
	fmt.Fprintf(writer, greeting)
}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	updatedTaskItems := append(taskItems, newTask)
	return updatedTaskItems
}
