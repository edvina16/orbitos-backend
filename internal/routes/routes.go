package routes

import (
	"github.com/edvina16/orbitos-backend/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(api *echo.Group, boardHandler *handlers.BoardHandler, stateHandler *handlers.StateHandler, taskHandler *handlers.TaskHandler) {
	// Board routes
	board := api.Group("/boards")
	board.GET("", boardHandler.ListBoards)
	board.GET("/:id", boardHandler.GetBoardByID)
	board.POST("", boardHandler.CreateBoard)
	board.POST("/:id", boardHandler.UpdateBoard)
	board.DELETE("/:id", boardHandler.DeleteBoard)

	// State routes (nested under boards)
	state := board.Group("/:board_id/states")
	state.GET("", stateHandler.ListStates)
	state.POST("", stateHandler.CreateState)
	state.POST("/:state_id", stateHandler.UpdateState)
	state.DELETE("/:state_id", stateHandler.DeleteState)

	// Task routes (nested under states)
	task := state.Group("/:state_id/tasks")
	task.GET("", stateHandler.ListTasksByState)
	task.POST("", stateHandler.CreateTaskInState)
	task.PUT("/:task_id", stateHandler.UpdateTaskState)

	// Top-level task routes
	api.GET("/tasks", taskHandler.ListTasks)
	api.POST("/tasks", taskHandler.CreateTask)
	api.POST("/tasks/:id", taskHandler.UpdateTask)
	api.DELETE("/tasks/:id", taskHandler.DeleteTask)
}
