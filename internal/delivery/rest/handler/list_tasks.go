package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-game-dev/rest-tm/internal/models/tm"
	"github.com/labstack/echo/v4"
)

func (t *TmHandler) ListTasks(c echo.Context) error {
	params := models.ListTasksParams{
		Order:  "asc",
		SortBy: "created_at",
	}

	if completed := c.QueryParam("completed"); completed != "" {
		parsed, err := strconv.ParseBool(completed)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid value for 'completed'. Must be true or false.",
			})
		}
		params.Completed = parsed
	}

	if order := c.QueryParam("order"); order != "" {
		if order != "asc" && order != "desc" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid value for 'order'. Must be 'asc' or 'desc'.",
			})
		}
		params.Order = order
	}

	if sortBy := c.QueryParam("sort_by"); sortBy != "" {
		params.SortBy = sortBy
	}

	tasks, err := t.TmSrv.ListTasks(context.Background(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to list tasks",
		})
	}

	return c.JSON(http.StatusOK, tasks)
}
