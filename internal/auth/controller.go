package auth

import (
	"rest-api/pkg/config"
	"time"

	"rest-api/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service Service
	cfg     *config.Config
}

func NewController(service Service, cfg *config.Config) *Controller {
	return &Controller{
		service: service,
		cfg:     cfg,
	}
}

func (ctrl *Controller) Login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid JSON format")
	}

	// Call service untuk login
	token, userResponse, err := ctrl.service.Login(req.Email, req.Password)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	// Set cookie dengan token
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   ctrl.cfg.NodeEnv == "production",
		SameSite: "Lax",
	})

	return response.Success(c, fiber.StatusOK, "Login successfully.", fiber.Map{
		"token": token,
		"user":  userResponse,
	})
}

func (ctrl *Controller) Register(c *fiber.Ctx) error {
	var req RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Call service untuk register
	userResponse, err := ctrl.service.Register(req.Username, req.Email, req.Password)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "User registered successfully.", fiber.Map{
		"user": userResponse,
	})
}
