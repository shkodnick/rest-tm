package models

import (
	"time"
)

type Task struct {
	Id        string    `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Body      string    `db:"body" json:"body"`
	Completed bool      `db:"completed" json:"completed"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CreateTaskParams struct {
	Title     string
	Body      string
	Completed bool
}

type Tasks struct {
	Tasks []Task
}

type ListTasksParams struct {
	Completed bool
	Order     string
	SortBy    string
}

type GetTaskResponse struct {
	Id        string
	Title     string
	Body      string
	Completed bool
}

type UpdateTaskResponse struct {
	Id        string
	Title     string
	Body      string
	Completed bool
	UpdatedAt  time.Time
}
