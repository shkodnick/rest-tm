package handler

import "github.com/labstack/echo/v4"

func (h *TmHandler) TestTask(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"message": "Task tested successfully",
	})
}
