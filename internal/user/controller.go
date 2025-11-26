package user

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

// @Summary Get user profile
// @Description Get current user profile
// @Tags User
// @Produce json
// @Success 200 {object} response.SuccessResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /api/users/profile [get]
func (ctrl *Controller) GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.User)

	userResponse, err := ctrl.service.GetProfile(user.ID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Profile retrieved successfully", fiber.Map{
		"user": userResponse,
	})
}

// @Summary Update user profile
// @Description Update user profile by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param data body UpdateRequest true "User data"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /api/users/{id} [put]
func (ctrl *Controller) UpdateUser(c *fiber.Ctx) error {
	currentUser := c.Locals("user").(*auth.User)
	id := c.Params("id")

	targetUserID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
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
		return response.Error(c, statusCode, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Profile updated successfully", fiber.Map{
		"user": userResponse,
	})
}

func (ctrl *Controller) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	userID, err := strconv.ParseUint(id, 10, 32)
	if (err != nil) {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	userResponse, err := ctrl.service.GetUserByID(uint(userID))
	if err != nil {
		statusCode := fiber.StatusInternalServerError
		if err.Error() == "user not found" {
			statusCode = fiber.StatusNotFound
		}
		return response.Error(c, statusCode, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "User retrieved successfully", fiber.Map{
		"user": userResponse,
	})
}