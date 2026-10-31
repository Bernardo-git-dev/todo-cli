package main

import (
	"fmt"
	"os"
)

func main() {
	tips := `
Task-CLI is a task manager written in Go by "Bernardo Tati"

Usage:
    task-cli [command] <args...>

Commands:
    add <title>                  Add a new task
    update <id> <title>          Update a task title
    delete <id>                  Delete a task by ID
    list                         List all tasks
    list <filter>                List tasks filtered by: todo, in_progress, done
    mark-todo <id>               Mark a task as todo
    mark-in-progress <id>        Mark a task as in progress
    mark-done <id>               Mark a task as done
`
	// Inicializa o arquivo data.json se necessário
	Start()

	args := os.Args

	// Se não passou nenhum comando, mostra a ajuda
	if len(args) < 2 {
		fmt.Println(tips)
		return
	}

	// O comando principal fica no índice 1
	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Error: Missing task title.\nUsage: task-cli add <title>")
			return
		}
		title := args[2]
		if err := Add(title); err != nil {
			fmt.Printf("Error adding task: %v\n", err)
			return
		}
		fmt.Printf("Task %q added successfully!\n", title)

	case "update":
		if len(args) < 4 {
			fmt.Println("Error: Missing ID or new title.\nUsage: task-cli update <id> <new_title>")
			return
		}
		id := args[2]
		newTitle := args[3]
		Update(id, newTitle)

	case "delete":
		if len(args) < 3 {
			fmt.Println("Error: Missing task ID.\nUsage: task-cli delete <id>")
			return
		}
		id := args[2]
		Delete(id)

	case "list":
		// Se passou um filtro opcional (ex: task-cli list done)
		if len(args) > 2 {
			filterInput := args[2]
			var filter Status

			switch filterInput {
			case "done":
				filter = StatusDone
			case "in-progress", "in_progress":
				filter = StatusInProgress
			case "todo":
				filter = StatusTodo
			default:
				fmt.Printf("Invalid filter %q. Choose between: todo, in_progress, done\n", filterInput)
				return
			}
			ListFilter(filter)
			return
		}
		List()

	case "mark-todo":
		if len(args) < 3 {
			fmt.Println("Error: Missing task ID.\nUsage: task-cli mark-todo <id>")
			return
		}
		Mark(StatusTodo, args[2])

	case "mark-in-progress", "mark-in-progres":
		if len(args) < 3 {
			fmt.Println("Error: Missing task ID.\nUsage: task-cli mark-in-progress <id>")
			return
		}
		Mark(StatusInProgress, args[2])

	case "mark-done":
		if len(args) < 3 {
			fmt.Println("Error: Missing task ID.\nUsage: task-cli mark-done <id>")
			return
		}
		Mark(StatusDone, args[2])

	default:
		fmt.Printf("Unknown command: %q\n", command)
		fmt.Println(tips)
	}
}
