package book

import (
	"fmt"

	"github.com/veyselaksin/library-api/utils"
)

func GetAllBooks() ([]map[string]interface{}, error) {
	var books utils.Book

	db, err := utils.ConnectDatabase()

	if err != nil {
		fmt.Println("Database connection error!")
	}
	result := []map[string]interface{}{}

	fmt.Println("Before the fill data: ", result)
	db.Model(&books).Find(&result)
	fmt.Println("After the fill data: ", result)

	return result, err
}

func GetSingleBook(id interface{}) (map[string]interface{}, error) {
	var book utils.Book

	db, err := utils.ConnectDatabase()

	if err != nil {
		fmt.Println("Database connection error!")
	}

	result := map[string]interface{}{}
	db.Model(&book).First(&result, id)

	return result, err
}
