package tm

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"github.com/go-game-dev/rest-tm/internal/models/tm"
	store "github.com/go-game-dev/rest-tm/internal/adapter/store"
)

type TmStore struct {
	Store *store.Store
}

func (t *TmStore) CreateTask(ctx context.Context, task models.CreateTaskParams) (models.Task, error) {

	result := models.Task{}

	sqlStr, args, err := squirrel.
		Insert(tableNameTask).
		Columns(fieldNameTitle, fieldNameBody, fieldNameCompleted).
		Values(task.Title, task.Body, task.Completed).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s, %s, %s",
			fieldNameTaskId,
			fieldNameTitle,
			fieldNameBody,
			fieldNameCompleted,
			fieldNameCreatedAt,
			fieldNameUpdatedAt)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Task{}, errors.Wrap(err, "ToSql")
	}

	err = t.Store.DB.Get(&result, sqlStr, args...)

	if err != nil {
		return models.Task{}, errors.Wrap(err, "Get")
	}

	return result, nil
}

func (t *TmStore) UpdateTask(ctx context.Context, task models.Task) (models.Task, error) {
	result := models.Task{}

	sqlStr, args, err := squirrel.
		Update(tableNameTask).
		Set(fieldNameTitle, task.Title).
		Set(fieldNameBody, task.Body).
		Set(fieldNameCompleted, task.Completed).
		Set(fieldNameUpdatedAt, squirrel.Expr("NOW()")).
		Where(squirrel.Eq{
			fieldNameTaskId: task.Id,
		}).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s, %s, %s",
			fieldNameTaskId,
			fieldNameTitle,
			fieldNameBody,
			fieldNameCompleted,
			fieldNameCreatedAt,
			fieldNameUpdatedAt)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Task{}, errors.Wrap(err, "ToSql")
	}

	err = t.Store.DB.QueryRow(sqlStr, args...).
		Scan(
			&result.Id,
			&result.Title,
			&result.Body,
			&result.Completed,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, errors.Wrap(err, "QueryRowContext")
	}

	return result, nil
}
func (t *TmStore) GetTask(ctx context.Context, taskId string) (models.Task, error) {
	result := models.Task{}

	query, args, err := squirrel.Select(
		fieldNameTaskId, fieldNameTitle, fieldNameBody, fieldNameCompleted,
	).
		From(tableNameTask).
		Where(squirrel.Eq{
			fieldNameTaskId: taskId,
		}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Task{}, errors.Wrap(err, "ToSql")
	}

	err = t.Store.DB.Get(&result, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Task{}, TaskNotFound
		}
		return models.Task{}, errors.Wrap(err, "Get")
	}

	return result, nil
}

func (t *TmStore) ListTasks(ctx context.Context, listTasksParams models.ListTasksParams) ([]models.Task, error) {
	result := make([]models.Task, 0)

	query := squirrel.Select(
		fieldNameTaskId, fieldNameTaskId, fieldNameBody, fieldNameCompleted, fieldNameCreatedAt, fieldNameUpdatedAt,
	).
		From(tableNameTask).
		PlaceholderFormat(squirrel.Dollar)
	
	if listTasksParams.Completed {
		query = query.Where(squirrel.Eq{"completed": listTasksParams.Completed})
	}

	query = query.OrderBy(fmt.Sprintf("%s %s", listTasksParams.SortBy, listTasksParams.Order))

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "ToSql")
	}

	err = t.Store.DB.SelectContext(ctx, &result, sqlStr, args...)
	if err != nil {
		return nil, errors.Wrap(err, "SelectContext")
	}

	return result, nil
}

func (t *TmStore) DeleteTask(ctx context.Context, taskID string) error {
	sqlStr, args, err := squirrel.
		Delete(tableNameTask).
		Where(squirrel.Eq{"task_id": taskID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "ToSql")
	}

	_, err = t.Store.DB.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		return errors.Wrap(err, "ExecContext")
	}

	return nil
}

