package tm

import (
	"context"

	"github.com/go-game-dev/rest-tm/internal/models/tm"
)

type TmRepository interface {
	CreateTask(ctx context.Context, task models.CreateTaskParams) (models.Task, error)
	UpdateTask(ctx context.Context, task models.Task) (models.Task, error)
	GetTask(ctx context.Context, task_id string) (models.Task, error)
	ListTasks(ctx context.Context, listTasksParams models.ListTasksParams) ([]models.Task, error)
	DeleteTask(ctx context.Context, task_id string) (error)
}
