package auth

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, ctrl *Controller) {
	auth := app.Group("/api/auth")
	
	auth.Post("/register", ctrl.Register)
	auth.Post("/login", ctrl.Login)
}