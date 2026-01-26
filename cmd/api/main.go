package main

import (
	"log"

	"github.com/edvina16/orbitos-backend/internal/database"
	"github.com/edvina16/orbitos-backend/internal/handlers"
	"github.com/edvina16/orbitos-backend/internal/routes"
	"github.com/edvina16/orbitos-backend/internal/service"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	elog "github.com/labstack/gommon/log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close db: %v", err)
		}
	}()

	queries := database.New(db)

	e := echo.New()
	e.Logger.SetLevel(elog.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	taskHandler := &handlers.TaskHandler{Service: &service.TaskService{DB: queries}}
	boardHandler := &handlers.BoardHandler{Service: &service.BoardService{DB: queries}}
	stateHandler := &handlers.StateHandler{Service: &service.StateService{DB: queries}, TaskService: &service.TaskService{DB: queries}}
	userHandler := &handlers.UserHandler{Service: &service.UserService{DB: queries}}

	e.POST("/api/register", userHandler.RegisterUser)
	e.POST("/api/login", userHandler.LoginUser)

	jwtSecret := "dev_secret" // Use env var in production
	apiGroup := e.Group("/api")
	apiGroup.Use(echojwt.JWT([]byte(jwtSecret)))

	routes.RegisterRoutes(apiGroup, boardHandler, stateHandler, taskHandler)

	e.Logger.Fatal(e.Start(":5000"))
}
