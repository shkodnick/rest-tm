package handler

import (
	"context"
	"net/http"

	"github.com/go-game-dev/rest-tm/internal/models/tm"
	"github.com/labstack/echo/v4"
)

func (t *TmHandler) UpdateTask(c echo.Context) error {

	taskID := c.Param("id")
	if taskID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Task ID is required",
		})
	}

	var params models.Task
	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	params.Id = taskID

	task, err := t.TmSrv.UpdateTask(context.Background(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update task",
		})
	}

	return c.JSON(http.StatusOK, task)
}