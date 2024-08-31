package router

import (
	"app/controllers"

	"github.com/gofiber/fiber/v2"
)

func User(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")

	user.Post("/signup", controllers.SignUp)
	user.Post("/login", controllers.LogIn)
	user.Get("/logout", controllers.LogOut)
	user.Put("/update-user/", controllers.UpdateUser)
	user.Put("/change-password", controllers.ChangePassword)
	user.Delete("/delete-account/", controllers.DeleteAccount)
	user.Get("/get-all-user", controllers.GetAllUser)
	user.Get("/:id", controllers.GetUserByID)
}
