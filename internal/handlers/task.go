package handlers

import (
	"github.com/edvina16/icpal-backend/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
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
