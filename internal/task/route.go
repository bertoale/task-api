package task

import (
	"rest-api/pkg/config"
	"rest-api/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, ctrl *Controller) {
	tasks := app.Group("/api/tasks")
	
	tasks.Post("/", middlewares.Auth(cfg), ctrl.CreateTask)
	tasks.Get("/", middlewares.Auth(cfg), ctrl.GetTasksByUserID)
	tasks.Get("/:id", middlewares.Auth(cfg), ctrl.GetTaskByID)
	tasks.Put("/:id", middlewares.Auth(cfg), ctrl.UpdateTask)
	tasks.Delete("/:id", middlewares.Auth(cfg), ctrl.DeleteTask)
}