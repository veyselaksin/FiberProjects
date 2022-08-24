package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/veyselaksin/library-api/pkg/book"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=/Turkey"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(err)
	if err != nil {
		fmt.Println("Failed to connect database")
	}
	fmt.Println("Database connection successfuly opened!")
}

func main() {
	app := fiber.New()
	initDatabase()

	fmt.Println(os.Getwd())
	setupRoutes(app)

	app.Listen(":3000")
}
