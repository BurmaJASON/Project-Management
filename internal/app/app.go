package app

import (
	"log"
	"net/http"
	"project-management/internal/project"
	"project-management/internal/task"
	"project-management/pkg/database"

	_ "project-management/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type App struct {
	Echo *echo.Echo
}

func NewApp() (*App, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	app := &App{
		Echo: echo.New(),
	}

	app.Echo.Use(middleware.Logger())
	app.Echo.Use(middleware.Recover())

	app.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	app.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	projectRepo := project.NewRepository(db)
	projectService := project.NewService(projectRepo)
	projectHandler := project.NewHandler(projectService)

	app.Echo.POST("/projects", projectHandler.CreateProject)
	app.Echo.GET("/projects/:id", projectHandler.GetProject)
	app.Echo.PUT("/projects/:id", projectHandler.UpdateProject)
	app.Echo.DELETE("/projects/:id", projectHandler.DeleteProject)

	taskRepo := task.NewRepository(db)
	taskService := task.NewService(taskRepo)
	taskHandler := task.NewHandler(taskService)

	app.Echo.GET("tasks/:id", taskHandler.GetTask)
	app.Echo.POST("/tasks", taskHandler.CreateTask)
	app.Echo.PUT("/tasks/:id", taskHandler.UpdateTask)
	app.Echo.DELETE("/tasks/:id", taskHandler.DeleteTask)

	return app, nil
}

func (a *App) Run() {
	log.Fatal(a.Echo.Start(":8080"))
}
