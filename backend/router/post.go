package router

import (
	"app/controllers"

	"github.com/gofiber/fiber/v2"
)

func Post(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	post := v1.Group("/post")

	post.Post("/", controllers.CreatePost)
	post.Get("/get-all/post", controllers.GetAllPost)
}
