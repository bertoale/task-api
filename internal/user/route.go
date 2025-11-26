package user

import (
	"rest-api/pkg/config"
	"rest-api/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, ctrl *Controller) {
	users := app.Group("/api/users")
	
	users.Get("/profile", middlewares.Auth(cfg), ctrl.GetProfile)
	users.Get("/:id", ctrl.GetUserByID)
	users.Put("/:id", middlewares.Auth(cfg), ctrl.UpdateUser)
}