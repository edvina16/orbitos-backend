package handlers

import (
	"net/http"
	"strconv"

	"github.com/edvina16/orbitos-backend/internal/service"
	"github.com/labstack/echo/v4"
)

type ReminderHandler struct {
	Service *service.ReminderService
	Task    *service.TaskService
}

func (h *ReminderHandler) CreateReminder(c echo.Context) error {
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	taskIDParam := c.Param("task_id")
	if taskIDParam == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing task_id in path"})
	}
	taskID, err := strconv.Atoi(taskIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task_id"})
	}
	var input struct {
		Message   string `json:"message"`
		RemindAt  string `json:"remind_at"`
		Frequency string `json:"frequency"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	if input.RemindAt == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing required fields"})
	}
	reminder, err := h.Service.CreateReminder(c.Request().Context(), userID, taskID, input.Message, input.RemindAt, input.Frequency)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, reminder)
}
