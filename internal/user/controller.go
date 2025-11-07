package user

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

func (ctrl *Controller) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	userResponse, err := ctrl.service.GetUserByID(uint(userID))
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if err.Error() == "user not found" {
			statusCode = fiber.StatusNotFound
		}
		return c.Status(statusCode).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User retrieved successfully",
		"user":    userResponse,
	})
}

func (ctrl *Controller) GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)

	userResponse, err := ctrl.service.GetProfile(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Profile retrieved successfully",
		"user":    userResponse,
	})
}

func (ctrl *Controller) UpdateUser(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*auth.User)
	id := c.Params("id")

	targetUserID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	userResponse, err := ctrl.service.UpdateUser(currentUser.ID, uint(targetUserID), &req)
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if err.Error() == "user not found" {
			statusCode = fiber.StatusNotFound
		} else if err.Error() == "unauthorized to update this user" {
			statusCode = fiber.StatusForbidden
		} else if err.Error() == "email already in use" || err.Error() == "username already in use" {
			statusCode = fiber.StatusBadRequest
		}
		return c.Status(statusCode).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Profile updated successfully",
		"user":    userResponse,
	})
}