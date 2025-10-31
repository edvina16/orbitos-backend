package handlers

import (
	"net/http"
	"strconv"

	"github.com/edvina16/icpal-backend/internal/service"
	"github.com/labstack/echo/v4"
)

type BoardHandler struct {
	Service *service.BoardService
}

func (h *BoardHandler) ListBoards(c echo.Context) error {
	boards, err := h.Service.ListBoards(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, boards)
}

func (h *BoardHandler) GetBoardByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid board ID"})
	}
	board, err := h.Service.GetBoardByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, board)
}

func (h *BoardHandler) CreateBoard(c echo.Context) error {
	var input struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	board, err := h.Service.CreateBoard(c.Request().Context(), input.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, board)
}
