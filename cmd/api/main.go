package main

import (
	"log"

	"github.com/edvina16/atmon-backend/internal/database"
	"github.com/edvina16/atmon-backend/internal/handlers"
	"github.com/edvina16/atmon-backend/internal/routes"
	"github.com/edvina16/atmon-backend/internal/service"
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
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	taskService := &service.TaskService{DB: queries}
	taskHandler := &handlers.TaskHandler{Service: taskService}
	routes.RegisterTask(e, taskHandler)

	boardService := &service.BoardService{DB: queries}
	boardHandler := &handlers.BoardHandler{Service: boardService}
	routes.RegisterBoard(e, boardHandler)

	stateService := &service.StateService{DB: queries}
	stateHandler := &handlers.StateHandler{Service: stateService, TaskService: taskService}
	routes.RegisterState(e, stateHandler)

	e.Logger.Fatal(e.Start(":5000"))
}
