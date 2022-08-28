package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/veyselaksin/library-api/utils"
)

func GetBooks(c *fiber.Ctx) error {
	books, err := GetAllBooks()

	if err != nil {
		return c.Status(404).JSON(
			utils.Response{
				Success: true,
				Message: "Not found!",
				Data:    []interface{}{},
			},
		)
	}
	return c.Status(200).JSON(
		utils.Response{
			Success: true,
			Message: "",
			Data:    books,
		},
	)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := GetSingleBook(id)

	if err != nil {
		return c.Status(404).JSON(
			utils.Response{
				Success: true,
				Message: "Not found!",
				Data:    []interface{}{},
			},
		)
	}

	return c.Status(200).JSON(
		utils.Response{
			Success: true,
			Message: "",
			Data:    book,
		},
	)
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("New book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete book")
}
