package handlers

import (
	"net/http"
	"strconv"

	"github.com/edvina16/orbitos-backend/internal/service"
	"github.com/labstack/echo/v4"
)

type BoardHandler struct {
	Service *service.BoardService
}

func (h *BoardHandler) ListBoards(c echo.Context) error {
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		c.Logger().Errorf("JWT extraction error: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	boards, err := h.Service.ListBoards(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, boards)
}

func (h *BoardHandler) GetBoardByID(c echo.Context) error {
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		c.Logger().Errorf("JWT extraction error: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid board ID"})
	}
	board, err := h.Service.GetBoardByID(c.Request().Context(), id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, board)
}

func (h *BoardHandler) CreateBoard(c echo.Context) error {
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		c.Logger().Errorf("JWT extraction error: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	c.Logger().Infof("CreateBoard using userID: %d", userID)
	var input struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	board, err := h.Service.CreateBoard(c.Request().Context(), input.Name, userID)
	if err != nil {
		c.Logger().Errorf("CreateBoard error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, board)
}

func (h *BoardHandler) DeleteBoard(c echo.Context) error {
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		c.Logger().Errorf("JWT extraction error: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Logger().Errorf("Invalid board ID: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid board ID"})
	}
	err = h.Service.DeleteBoard(c.Request().Context(), id, userID)
	if err != nil {
		c.Logger().Errorf("Failed to delete board %d: %v", id, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]bool{"success": true})
}

func (h *BoardHandler) UpdateBoard(c echo.Context) error {
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		c.Logger().Errorf("JWT extraction error: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid board ID"})
	}
	var input struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	err = h.Service.UpdateBoard(c.Request().Context(), id, input.Name, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	board, err := h.Service.GetBoardByID(c.Request().Context(), id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, board)
}
