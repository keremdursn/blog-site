package main

import (
	"app/database"
	"app/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	router.User(app)
	router.Post(app)
	app.Listen(":8080")
}
