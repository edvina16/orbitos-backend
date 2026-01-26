package handlers

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetUserIDFromContext(c echo.Context) (int, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return 0, errors.New("missing or invalid JWT token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid JWT claims")
	}
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user_id in JWT claims")
	}
	return int(userIDFloat), nil
}
