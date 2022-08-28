package utils

import "gorm.io/gorm"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Book struct {
	gorm.Model
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Rating float64 `json:"rating"`
}
