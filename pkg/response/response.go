package response

import "github.com/gofiber/fiber/v2"

// SuccessResponse is a standard success response for Swagger docs
// @Description Success response
// @example {"success":true,"message":"string","data":{}}
type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Operation successful"`
	Data    interface{} `json:"data"`
}

// ErrorResponse is a standard error response for Swagger docs
// @Description Error response
// @example {"success":false,"message":"error message"}
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"error message"`
}

func Success(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func Error(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}
