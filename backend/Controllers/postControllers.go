package controllers

import (
	"app/database"
	"app/model"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "data": ok})
	}

	db := database.DB.Db
	post := new(model.Post)
	c.BodyParser(&post)

	posttitle := c.FormValue("posttitle")
	postdesc := c.FormValue("postdesc")
	post.UserID = user.ID

	if posttitle != "" {
		post.Title = posttitle
	}
	if postdesc != "" {
		post.Description = postdesc
	}

	err := db.Create(&post).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error"})
	}

	err = database.DB.Db.Preload("User").First(&post, post.ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "error to innerjoin", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "post created", "data": post})
}

//update post
//delete post
//like post
//get all
//get by id
