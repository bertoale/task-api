package routes

import (
	"rest-api/config"
	"rest-api/internal/auth"
	"rest-api/internal/database"
	"rest-api/internal/task"
	"rest-api/internal/user"

	"github.com/gofiber/fiber/v2"
)

// SetupVerticalRoutes sets up routes using the vertical layer architecture
func SetupVerticalRoutes(app *fiber.App, cfg *config.Config) {
	db := database.GetDB()

	// Initialize Auth module (vertical)
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, cfg)
	authController := auth.NewController(authService, cfg)
	auth.SetupRoutes(app, authController)

	// Initialize User module (vertical)
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userController := user.NewController(userService)
	user.SetupRoutes(app, cfg, userController)

	// Initialize Task module (vertical)
	taskRepo := task.NewRepository(db)
	taskService := task.NewService(taskRepo)
	taskController := task.NewController(taskService)
	task.SetupRoutes(app, cfg, taskController)
}
