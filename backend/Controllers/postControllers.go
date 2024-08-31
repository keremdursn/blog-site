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
	posts := new(model.Post)
	c.BodyParser(&posts)

	poststitle := c.FormValue("poststitle")
	postsdesc := c.FormValue("postsdesc")
	posts.UserID = user.ID

	if poststitle != "" {
		posts.Title = poststitle
	}
	if postsdesc != "" {
		posts.Description = postsdesc
	}

	err := db.Create(&posts).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error"})
	}

	err = database.DB.Db.Preload("User").First(&posts, posts.ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "error to innerjoin", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "post created", "data": posts})
}

// func UpdatePost(c *fiber.Ctx) error {
// 	user, ok := c.Locals("user").(model.User)
// 	if !ok {
// 		return c.Status(400).JSON(fiber.Map{"Status": "Error", "data": ok})
// 	}

// 	db := database.DB.Db
// 	posts := new(model.Post)
// 	c.BodyParser(posts)
// 	err := db.Where("user_id = ?", user.ID).First(&posts).Error
// 	if err != nil {
// 		return  c.Status(404).JSON(fiber.Map{"status": "error", "message": ""})
// 	}

// }

// func DeletePost(c *fiber.Ctx) error {
// 	user, ok := c.Locals("user").(model.User)
// 	if !ok {
// 		return c.Status(400).JSON(fiber.Map{"Status": "Error", "data": ok})
// 	}
// 	db := database.DB.Db
// 	posts := new(model.Post)
// 	c.BodyParser(posts)

// }

func GetAllPost(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(model.User)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "data": ok})
	}
	db := database.DB.Db
	posts := new([]model.Post)

	err := db.Preload("User").Where("user_id = ?", user.ID).Order("id DESC").Find(posts).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "failed", "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "data": posts})
}

//update post
//delete post
//like post
//get all
//get by id
