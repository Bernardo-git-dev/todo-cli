package main

import (
	"os"
)

func main() {
	Start()

	args := os.Args

	switch args[1] {
	case "list":
		List()
	case "add":
		title := os.Args[2]
		Add(title)
	case "update":
		id := os.Args[2]
		newTitle := os.Args[3]
		Update(id, newTitle)
	case "delete":
		id := os.Args[2]
		Delete(id)
	}
}
