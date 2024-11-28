package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/go-game-dev/rest-tm/internal/bootstrap"
	"github.com/go-game-dev/rest-tm/internal/delivery/rest/handler"
)

func RegisterRoutes(e *echo.Echo) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tmSrv, err := bootstrap.CreateTmService()
	if err != nil {
		return err
	}

	tmHandler := handler.NewTmHandler(tmSrv)
	
	t := e.Group("/task")
	t.POST("/create", tmHandler.CreateTask)
	t.PUT("/update/:id", tmHandler.UpdateTask)
	t.POST("/list", tmHandler.ListTasks)
	t.GET("/get/:id", tmHandler.GetTask)
	t.DELETE("/delete/:id", tmHandler.DeleteTask)

	e.GET("/test",tmHandler.TestTask)

	return nil
}
