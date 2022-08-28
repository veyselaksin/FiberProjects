package train

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeService(app fiber.Router) {
	api := app.Group("/train")

	api.Get("/type/struct", TrainStruct)
}
