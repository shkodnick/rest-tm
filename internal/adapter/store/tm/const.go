package tm

import "github.com/pkg/errors"

const (
	tableNameTask = "task"

	fieldNameTaskId    = "task_id"
	fieldNameTitle     = "title"
	fieldNameBody      = "body"
	fieldNameCompleted = "completed"
	fieldNameCreatedAt = "created_at"
	fieldNameUpdatedAt = "updated_at"
)

var TaskNotFound = errors.New("this task not found")
