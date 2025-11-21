package main

import (
	"library_management2/controllers"
	"library_management2/services"
)

func main() {
	// 1. Create a new library service (empty books & members)
	library := services.NewLibrary()

	// 2. Create a controller, attach the library service
	controller := controllers.NewLibraryController(library)

	// 3. Start the console program
	controller.Start()
}
