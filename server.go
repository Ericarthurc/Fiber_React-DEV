package main

import (
	"fmt"

	"ericarthurc/fiberAPI/database"
	"ericarthurc/fiberAPI/router"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// Connect to database
	database.ConnectDB()

	app := fiber.New(&fiber.Settings{
		// Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
	})

	// Middleware
	app.Use(middleware.Logger())
	app.Use(cors.New())

	router.UserRoutes(app)

	app.Static("/", "./frontend/build")

	fmt.Println("Server running on port 5010")
	app.Listen(5010)

	defer database.DB.Close()
}
