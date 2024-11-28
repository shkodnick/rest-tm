package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (t *TmHandler) DeleteTask(c echo.Context) error {
	taskID := c.Param("id")
	if taskID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Task ID is required",
		})
	}

	err := t.TmSrv.DeleteTask(context.Background(), taskID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete task",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Task deleted successfully",
	})
}
