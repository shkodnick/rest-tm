package handler

import (
	"context"
	"net/http"

	"github.com/go-game-dev/rest-tm/internal/models/tm"
	"github.com/labstack/echo/v4"
)

func (t *TmHandler) ListTasks(c echo.Context) error {
	var params models.ListTasksParams

	if err := c.Bind(&params); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "Invalid request body",
			})
	}

	tasks, err := t.TmSrv.ListTasks(context.Background(), params)
	if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Failed to retrieve tasks",
			})
	}

	return c.JSON(http.StatusOK, tasks)
}

