package auth

import (
	"rest-api/config"
	"time"

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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
		})
	}

	// Call service untuk login
	token, userResponse, err := ctrl.service.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
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

	return c.JSON(fiber.Map{
		"message": "Login successfully.",
		"token":   token,
		"user":    userResponse,
	})
}

func (ctrl *Controller) Register(c *fiber.Ctx) error {
	var req RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
		})
	}

	// Call service untuk register
	userResponse, err := ctrl.service.Register(req.Username, req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully.",
		"user":    userResponse,
	})
}
