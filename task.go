package main

import "time"

type Status string

const (
	StatusTodo      Status = "todo"
	StatusInprogres Status = "in_progress"
	StatusDone      Status = "done"
)

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
