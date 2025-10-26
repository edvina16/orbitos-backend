package routes

import (
	"github.com/edvina16/icpal-backend/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterTask(e *echo.Echo, taskHandler *handlers.TaskHandler) {
	e.GET("/api/tasks", taskHandler.ListTasks)
	e.POST("/api/tasks", taskHandler.CreateTask)
}

func RegisterBoard(e *echo.Echo, boardHandler *handlers.BoardHandler) {
	e.GET("/api/boards", boardHandler.ListBoards)
	e.GET("/api/boards/:id", boardHandler.GetBoardByID)
	e.POST("/api/boards", boardHandler.CreateBoard)
}

func RegisterState(e *echo.Echo, stateHandler *handlers.StateHandler) {
	e.GET("/api/boards/:board_id/states", stateHandler.ListStates)
}
