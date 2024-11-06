package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var tasks []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTo-Do List Manager")
		fmt.Println("1. View tasks")
		fmt.Println("2. Add a task")
		fmt.Println("3. Delete a task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			viewTasks()
		case "2":
			addTask(scanner)
		case "3":
			deleteTask(scanner)
		case "4":
			fmt.Println("Exiting To-Do List Manager.")
			return
		default:
			fmt.Println("Invalid option. Please choose 1, 2, 3, or 4.")
		}
	}
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks in the list.")
	} else {
		fmt.Println("Tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter the task: ")
	scanner.Scan()
	task := scanner.Text()
	tasks = append(tasks, task)
	fmt.Println("Task added successfully.")
}

func deleteTask(scanner *bufio.Scanner) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to delete.")
		return
	}

	fmt.Print("Enter the task number to delete: ")
	scanner.Scan()
	taskNumStr := scanner.Text()
	taskNum, err := strconv.Atoi(taskNumStr)

	if err != nil || taskNum < 1 || taskNum > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)
	fmt.Println("Task deleted successfully.")
}
