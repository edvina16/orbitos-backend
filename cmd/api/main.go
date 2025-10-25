package main

import (
	"log"

	"github.com/edvina16/icpal-backend/internal/database"
	"github.com/edvina16/icpal-backend/internal/handlers"
	"github.com/edvina16/icpal-backend/internal/routes"
	"github.com/edvina16/icpal-backend/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	queries := database.New(db)

	e := echo.New()
	e.Use(middleware.CORS())

	taskService := &service.TaskService{DB: queries}
	taskHandler := &handlers.TaskHandler{Service: taskService}
	routes.Register(e, taskHandler)

	e.Logger.Fatal(e.Start(":5000"))
}
