package main

import (
	"fmt"

	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	library := services.NewLibrary(
		[]models.Book{
			{ID: 1, Title: "1984", Author: "George Orwell", Status: "Available"},
			{ID: 2, Title: "The Pragmatic Programmer", Author: "Andrew Hunt", Status: "Available"},
			{ID: 3, Title: "Clean Code", Author: "Robert C. Martin", Status: "Available"},
		},
		[]models.Member{
			{ID: 1, Name: "Alice Johnson"},
			{ID: 2, Name: "Bob Smith"},
		},
	)

	controller := controllers.NewLibraryController(library)

	fmt.Println("Welcome to the Console Library Management System")
	controller.Run()
}
