package algorithm

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/veyselaksin/kwin-api/utils"
)

func BinarySearch(c *fiber.Ctx) error {
	fmt.Println("gfsd")

	return c.Status(200).JSON(
		utils.Response{
			Success: true,
			Message: "Not found!",
			Data:    []interface{}{},
		},
	)
}
