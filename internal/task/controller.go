package task

import (
	"rest-api/internal/auth"
	"rest-api/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

// @Summary Create new task
// @Description Buat task baru untuk user
// @Tags Tasks
// @Accept json
// @Produce json
// @Param data body CreateRequest true "Task data"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /api/tasks [post]
func (ctrl *Controller) CreateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)

	var req CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	taskResponse, err := ctrl.service.CreateTask(user.ID, req.Title, req.Description)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "Task created successfully", fiber.Map{
		"task": taskResponse,
	})
}

// @Summary List user tasks
// @Description Get all tasks for current user
// @Tags Tasks
// @Produce json
// @Success 200 {object} response.SuccessResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /api/tasks [get]
func (ctrl *Controller) GetTasksByUserID(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)

	tasks, err := ctrl.service.GetTasksByUserID(user.ID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Tasks retrieved successfully", fiber.Map{
		"tasks": tasks,
	})
}

// @Summary Get task detail
// @Description Get detail of a task by ID
// @Tags Tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /api/tasks/{id} [get]
func (ctrl *Controller) GetTaskByID(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)
	id := c.Params("id")

	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid task ID")
	}

	task, err := ctrl.service.GetTaskByID(user.ID, uint(taskID))
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if err.Error() == "task not found" {
			statusCode = fiber.StatusNotFound
		} else if err.Error() == "unauthorized to access this task" {
			statusCode = fiber.StatusForbidden
		}
		return response.Error(c, statusCode, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Task retrieved successfully", fiber.Map{
		"task": task,
	})
}

// @Summary Update task
// @Description Update a task by ID
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param data body UpdateRequest true "Task data"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /api/tasks/{id} [put]
func (ctrl *Controller) UpdateTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)
	id := c.Params("id")

	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid task ID")
	}

	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
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
		return response.Error(c, statusCode, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Task updated successfully", fiber.Map{
		"task": updatedTask,
	})
}

// @Summary Delete task
// @Description Delete a task by ID
// @Tags Tasks
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /api/tasks/{id} [delete]
func (ctrl *Controller) DeleteTask(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)
	id := c.Params("id")

	taskID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid task ID")
	}

	if err := ctrl.service.DeleteTask(user.ID, uint(taskID)); err != nil {
		statusCode := fiber.StatusInternalServerError
		if err.Error() == "task not found" {
			statusCode = fiber.StatusNotFound
		} else if err.Error() == "unauthorized to delete this task" {
			statusCode = fiber.StatusForbidden
		}
		return response.Error(c, statusCode, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Task deleted successfully", fiber.Map{})
}