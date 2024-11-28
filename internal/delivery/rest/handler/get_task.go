package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (t *TmHandler) GetTask(c echo.Context) error {
	taskID := c.Param("id")
	if taskID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Task ID is required",
		})
	}

	task, err := t.TmSrv.GetTask(context.Background(), taskID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to get task",
		})
	}

	return c.JSON(http.StatusCreated, task)
}