package handlers

import (
	"github.com/edvina16/icpal-backend/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type StateHandler struct {
	Service *service.StateService
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
