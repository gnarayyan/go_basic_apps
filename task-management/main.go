package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gnarayyan/go_basic_apps/task-management/storage"
	"github.com/gnarayyan/go_basic_apps/task-management/task"
)

func displayMenu() {
	fmt.Println("\n--- Task Management Menu ---")
	fmt.Println("1. View all tasks")
	fmt.Println("2. View a single task")
	fmt.Println("3. Add a new task")
	fmt.Println("4. Delete a task")
	fmt.Println("5. Exit")
	fmt.Print("Choose an option (1-5): ")
}

func main() {
	store := storage.NewMemoryStorage()

	// Initial tasks
	store.AddTask(&task.Task{Desc: "Learn Go", Status: "pending", CreatedAt: time.Now()})
	store.AddTask(&task.Task{Desc: "Write code", Status: "completed", CreatedAt: time.Now()})

	for {
		displayMenu()

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || choice < 1 || choice > 5 {
			fmt.Println("Invalid input, please choose a number between 1 and 5.")
			continue
		}

		switch choice {
		case 1:
			// View all tasks
			tasks := store.ListTasks()
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")
			} else {
				fmt.Println("\nAll tasks:")
				for _, t := range tasks {
					fmt.Printf("Task %d: %s (%s)\n", t.ID, t.Desc, t.Status)
				}
			}

		case 2:
			// View a single task
			fmt.Print("Enter the task ID to view: ")
			var taskID int
			_, err := fmt.Scan(&taskID)
			if err != nil {
				fmt.Println("Invalid task ID.")
				continue
			}
			task, exists := store.GetTask(taskID)
			if !exists {
				fmt.Println("Task not found.")
			} else {
				fmt.Printf("\nTask %d: %s (%s)\n", task.ID, task.Desc, task.Status)
			}

		case 3:
			// Add a new task
			fmt.Print("Enter task description: ")
			var desc string

			fmt.Scan(&desc)
			status := "pending"
			task := &task.Task{Desc: desc, Status: status, CreatedAt: time.Now()}
			store.AddTask(task)
			fmt.Println("New task added.")

		case 4:
			// Delete a task
			fmt.Print("Enter the task ID to delete: ")
			var taskID int
			_, err := fmt.Scan(&taskID)
			if err != nil {
				fmt.Println("Invalid task ID.")
				continue
			}
			if store.DeleteTask(taskID) {
				fmt.Println("Task deleted successfully.")
			} else {
				fmt.Println("Task not found.")
			}

		case 5:
			// Exit
			fmt.Println("Exiting Task Management System.")
			os.Exit(0)
		}
	}
}
