package utils

import "github.com/gofiber/fiber/v2"

func ErrorResponse(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(
		Response{
			Message: msg,
			Data:    []interface{}{},
			Success: false,
		})
}

func SuccessResponse(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(
		Response{
			Message: "success",
			Data:    data,
			Success: true,
		})
}
