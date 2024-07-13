package main

import (
	"fmt"
	"strconv"
)


type Todo struct {
	Id int 
	Description string
	Complete bool
}

var todos []Todo


func main() {
	for {
		DisplayMenu()
		choice := getUserInput("Choose an option")
		switch choice {
		case "1":
			addTask()
		case "2":
			listTasks()
		case "3":
			completeTask()
		case "4":
			deleteTask()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice: please try again")
		}
	}
}

func DisplayMenu () {
	fmt.Println("1. Add Task")
    fmt.Println("2. List Tasks")
    fmt.Println("3. Complete Task")
    fmt.Println("4. Delete Task")
    fmt.Println("5. Exit")

}

func getUserInput(prompt string) string {
    var input string
    fmt.Print(prompt + ": ")
    fmt.Scanln(&input)
    return input
}

func addTask() {
	description := getUserInput("Enter task description")
	task := Todo{Id: len(todos) + 1, Description: description, Complete: false}
	todos = append(todos, task)
	fmt.Println("****************************************************************")
}


func listTasks() {
	
	fmt.Println("****************************************************************")
	if len(todos) == 0 {
		fmt.Println("its empty")
	}
	for _, task := range todos {
		status := "Pending"
		if task.Complete {
			status = "Completed"
		}
		fmt.Printf("%d. %s [%s]\n", task.Id, task.Description, status)
	}
	fmt.Println("****************************************************************")
}

func completeTask () {
	idStr := getUserInput("Enter the ID of the task to mark as complete")
	id , _ :=strconv.Atoi(idStr)
	for i , task := range todos {
		if task.Id == id {
			todos[i].Complete = true
			break
		}
	}
}

func deleteTask() {
	idStr := getUserInput("Enter the ID of the task to delete")
	id , _ :=strconv.Atoi(idStr)
	for i , task := range todos {
		if task.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
}