package routes

import (
	"github.com/edvina16/atmon-backend/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterTask(e *echo.Echo, taskHandler *handlers.TaskHandler) {
	e.GET("/api/tasks", taskHandler.ListTasks)
	e.POST("/api/tasks", taskHandler.CreateTask)
	e.POST("/api/tasks/:id", taskHandler.UpdateTask)
	e.DELETE("/api/tasks/:id", taskHandler.DeleteTask)
}

func RegisterBoard(e *echo.Echo, boardHandler *handlers.BoardHandler) {
	e.GET("/api/boards", boardHandler.ListBoards)
	e.GET("/api/boards/:id", boardHandler.GetBoardByID)
	e.POST("/api/boards", boardHandler.CreateBoard)
	e.POST("/api/boards/:id", boardHandler.UpdateBoard)
	e.DELETE("/api/boards/:id", boardHandler.DeleteBoard)
}

func RegisterState(e *echo.Echo, stateHandler *handlers.StateHandler) {
	e.GET("/api/boards/:board_id/states", stateHandler.ListStates)
	e.POST("/api/boards/:board_id/states", stateHandler.CreateState)
	e.POST("/api/boards/:board_id/states/:state_id", stateHandler.UpdateState)
	e.DELETE("/api/boards/:board_id/states/:state_id", stateHandler.DeleteState)

	// Task-related state routes
	e.GET("/api/boards/:board_id/states/:state_id/tasks", stateHandler.ListTasksByState)
	e.POST("/api/boards/:board_id/states/:state_id/tasks", stateHandler.CreateTaskInState)
	e.PUT("/api/boards/:board_id/states/:state_id/tasks/:task_id", stateHandler.UpdateTaskState)
}
