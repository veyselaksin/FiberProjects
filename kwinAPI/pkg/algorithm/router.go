package algorithm

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeService(app fiber.Router) {
	api := app.Group("/algorithm")

	api.Get("/binary-search", BinarySearch)
}
