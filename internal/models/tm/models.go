package models

import "time"

type Task struct {
	Id        string
	Title     string
	Body      string
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
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

