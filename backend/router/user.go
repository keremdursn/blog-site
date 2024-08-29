package router

import (
	"app/controllers"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	v1.Post("/signup", controllers.SignUp)
	v1.Get("/login", controllers.LogIn)
}
