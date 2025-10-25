package routes

import (
	"github.com/edvina16/icpal-backend/internal/handlers"
	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, taskHandler *handlers.TaskHandler) {
	e.GET("/api/tasks", taskHandler.ListTasks)
	e.POST("/api/tasks", taskHandler.CreateTask)
}
