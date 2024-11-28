package tm

import (
	"context"

	"github.com/go-game-dev/rest-tm/internal/models/tm"
)

type TmRepository interface {
	CreateTask(ctx context.Context, task models.CreateTaskParams) (models.Task, error)
	UpdateTask(ctx context.Context, task models.Task) (models.UpdateTaskResponse, error)
	GetTask(ctx context.Context, task_id string) (models.GetTaskResponse, error)
	ListTasks(ctx context.Context, listTasksParams models.ListTasksParams) ([]models.GetTaskResponse, error)
	DeleteTask(ctx context.Context, task_id string) (error)
}
