package router

import (
	"app/controllers"
	"app/middleware"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")

	user.Post("/signup", controllers.SignUp)
	user.Get("/login", controllers.LogIn)
	user.Put("/update-user/:id", middleware.TokenControl, controllers.UpdateUser)
	user.Delete("/delete-account/:id", middleware.TokenControl, controllers.DeleteAccount)
	user.Get("/get-all-user", controllers.GetAllUser)
	user.Get("/:id", controllers.GetUserByID)
}
