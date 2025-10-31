package handlers

import (
	"net/http"
	"strconv"

	"github.com/edvina16/protai-backend/internal/service"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	Service *service.TaskService
}

func (h *TaskHandler) ListTasks(c echo.Context) error {
	tasks, err := h.Service.ListTasks(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	task, err := h.Service.CreateTask(c.Request().Context(), input.Title, input.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	taskIDStr := c.Param("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.Logger().Errorf("Invalid task ID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task ID"})
	}
	err = h.Service.DeleteTask(c.Request().Context(), taskID)
	if err != nil {
		c.Logger().Errorf("Failed to delete task %d: %v", taskID, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	taskIDStr := c.Param("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task ID"})
	}
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	err = h.Service.UpdateTask(c.Request().Context(), taskID, input.Title, input.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	task, err := h.Service.GetTaskByID(c.Request().Context(), int32(taskID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, task)
}
