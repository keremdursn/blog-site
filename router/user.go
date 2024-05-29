package router

import (
	"app/controllers"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App) {
	api := app.Group("/user")

	api.Post("/signup", controllers.SignUp)
	api.Get("/login", controllers.LogIn)
}
