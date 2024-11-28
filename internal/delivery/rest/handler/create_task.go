package handler

import (
	"context"
	"net/http"

	"github.com/go-game-dev/rest-tm/internal/models/tm"
	"github.com/labstack/echo/v4"
)

func (t *TmHandler) CreateTask(c echo.Context) error {
	var params models.CreateTaskParams

	if err := c.Bind(&params); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if params.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Title is required",
		})
	}

	task, err := t.TmSrv.CreateTask(context.Background(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create task",
		})
	}

	return c.JSON(http.StatusCreated, task)
}