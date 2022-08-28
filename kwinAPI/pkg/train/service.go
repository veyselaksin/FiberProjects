package train

import (
	"github.com/gofiber/fiber/v2"
	"github.com/veyselaksin/kwin-api/utils"
)

func TrainStruct(c *fiber.Ctx) error {
	return c.Status(200).JSON(utils.Response{
		Success: true,
		Message: "",
		Data:    []interface{}{},
	})
}
