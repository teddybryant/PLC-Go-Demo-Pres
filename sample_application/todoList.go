package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	todoList := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n To Do List:")
		fmt.Println("1. Add a task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Remove a task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			scanner.Scan()
			task := scanner.Text()
			todoList = append(todoList, task)
			fmt.Println("Task added.")

		case "2":
			fmt.Println("To-Do List:")
			for i, task := range todoList {
				fmt.Printf("%d. %s\n", i+1, task)
			}

		case "3":
			if len(todoList) <= 0 {
				fmt.Println("No tasks to remove.")
				continue
			}
			fmt.Print("Enter the task number to remove: ")
			scanner.Scan()
			input := scanner.Text()
			taskNum, err := strconv.Atoi(input)
			if err != nil || taskNum < 1 || taskNum > len(todoList) {
				fmt.Println("Invalid task number.")
				continue
			}
			taskIndex := taskNum - 1
			todoList = append(todoList[:taskIndex], todoList[taskIndex+1:]...)
			fmt.Println("Task removed.")

		case "4":
			fmt.Println("Exiting")
			return

		default:
			fmt.Println("Invalid option. Try something else")
		}
	}
}
