package task

import (
	"rest-api/internal/auth"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) CreateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)

	var req CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	taskResponse, err := ctrl.service.CreateTask(user.ID, req.Title, req.Description)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Task created successfully",
		"task":    taskResponse,
	})
}

func (ctrl *Controller) GetTasksByUserID(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)

	tasks, err := ctrl.service.GetTasksByUserID(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Tasks retrieved successfully",
		"tasks":   tasks,
	})
}

func (ctrl *Controller) GetTaskByID(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)
	id := c.Params("id")

	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid task ID",
		})
	}

	task, err := ctrl.service.GetTaskByID(user.ID, uint(taskID))
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if err.Error() == "task not found" {
			statusCode = fiber.StatusNotFound
		} else if err.Error() == "unauthorized to access this task" {
			statusCode = fiber.StatusForbidden
		}
		return c.Status(statusCode).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task retrieved successfully",
		"task":    task,
	})
}

func (ctrl *Controller) UpdateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)
	id := c.Params("id")

	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid task ID",
		})
	}

	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	updatedTask, err := ctrl.service.UpdateTask(user.ID, uint(taskID), &req)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if err.Error() == "task not found" {
			statusCode = fiber.StatusNotFound
		} else if err.Error() == "unauthorized to update this task" {
			statusCode = fiber.StatusForbidden
		} else if err.Error() == "title is required" {
			statusCode = fiber.StatusBadRequest
		}
		return c.Status(statusCode).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task updated successfully",
		"task":    updatedTask,
	})
}

func (ctrl *Controller) DeleteTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)
	id := c.Params("id")

	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid task ID",
		})
	}

	if err := ctrl.service.DeleteTask(user.ID, uint(taskID)); err != nil {
		statusCode := fiber.StatusInternalServerError
		if err.Error() == "task not found" {
			statusCode = fiber.StatusNotFound
		} else if err.Error() == "unauthorized to delete this task" {
			statusCode = fiber.StatusForbidden
		}
		return c.Status(statusCode).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}