package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/veyselaksin/kwin-api/pkg/algorithm"
	"github.com/veyselaksin/kwin-api/pkg/train"
)

func initializeRouters(app fiber.Router) {
	api := app.Group("api/v1/kwin")
	algorithm.InitializeService(api)
	train.InitializeService(api)

}

func main() {
	app := fiber.New()
	initializeRouters(app)

	app.Listen(":3000")
}
