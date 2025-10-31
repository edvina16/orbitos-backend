package handlers

import (
	"net/http"
	"strconv"

	"github.com/edvina16/atmon-backend/internal/service"
	"github.com/labstack/echo/v4"
)

type StateHandler struct {
	Service     *service.StateService
	TaskService *service.TaskService
}

func (h *StateHandler) ListStates(c echo.Context) error {
	boardIDStr := c.Param("board_id")
	boardID, err := strconv.Atoi(boardIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid board ID"})
	}
	states, err := h.Service.ListStatesByBoard(c.Request().Context(), boardID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, states)
}

func (h *StateHandler) CreateState(c echo.Context) error {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	boardIDStr := c.Param("board_id")
	boardID, err := strconv.Atoi(boardIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid board ID"})
	}

	state, err := h.Service.CreateState(c.Request().Context(), input.Name, boardID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, state)
}

func (h *StateHandler) ListTasksByState(c echo.Context) error {
	stateIDStr := c.Param("state_id")
	stateID, err := strconv.Atoi(stateIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid state ID"})
	}
	tasks, err := h.Service.ListTasksByState(c.Request().Context(), stateID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *StateHandler) CreateTaskInState(c echo.Context) error {
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	stateIDStr := c.Param("state_id")
	stateID, err := strconv.Atoi(stateIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid state ID"})
	}
	task, err := h.Service.CreateTaskInState(c.Request().Context(), input.Title, input.Content, stateID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, task)
}

func (h *StateHandler) UpdateTaskState(c echo.Context) error {
	taskIDStr := c.Param("task_id")
	stateIDStr := c.Param("state_id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task ID"})
	}
	stateID, err := strconv.Atoi(stateIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid state ID"})
	}
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	// Update all fields: state, title, content
	err = h.TaskService.UpdateTask(c.Request().Context(), taskID, input.Title, input.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	err = h.Service.UpdateTaskState(c.Request().Context(), taskID, stateID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	task, err := h.TaskService.GetTaskByID(c.Request().Context(), int32(taskID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *StateHandler) DeleteState(c echo.Context) error {
	stateIDStr := c.Param("state_id")
	stateID, err := strconv.Atoi(stateIDStr)
	if err != nil {
		c.Logger().Errorf("Invalid state ID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid state ID"})
	}
	err = h.Service.DeleteState(c.Request().Context(), stateID)
	if err != nil {
		c.Logger().Errorf("Failed to delete state %d: %v", stateID, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func (h *StateHandler) UpdateState(c echo.Context) error {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	stateIDStr := c.Param("state_id")
	stateID, err := strconv.Atoi(stateIDStr)
	boardIDStr := c.Param("board_id")
	boardID, err := strconv.Atoi(boardIDStr)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid state ID"})
	}
	err = h.Service.UpdateState(c.Request().Context(), stateID, input.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	state, err := h.Service.GetStateByID(c.Request().Context(), int32(stateID), int32(boardID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, state)
}
