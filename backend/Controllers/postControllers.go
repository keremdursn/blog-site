package controllers

import (
	"app/database"
	"app/middleware"
	"app/model"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server errorrrrr"})
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

	err = db.Create(&posts).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error"})
	}

	err = database.DB.Db.Preload("User").First(&posts, posts.ID).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "error to innerjoin", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "post created", "data": posts})
}

func UpdatePost(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server errorrrrr"})
	}
	id := c.Params("id")

	db := database.DB.Db
	posts := new(model.Post)
	// if err := c.BodyParser(&posts); err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"error": "Error parsing request"})
	// }

	err = db.First(&posts, id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Post not found"})
	}

	if user.ID != posts.UserID {
		return c.Status(401).JSON(fiber.Map{"status": "failed", "message": "you don't have permission update this post"})
	}

	poststitle := c.FormValue("poststitle")
	postsdesc := c.FormValue("postsdesc")

	if poststitle != "" {
		posts.Title = poststitle
	}
	if postsdesc != "" {
		posts.Description = postsdesc
	}
	if err := database.DB.Db.Save(&posts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"Status": "Error", "Message": "cant update"})
	}

	return c.Status(200).JSON(fiber.Map{"Status": "Success", "Message": "Success", "data": posts})

}

 func DeletePost(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server errorrrrr"})
	}

 	db := database.DB.Db
 	posts := new(model.Post)
	
	id := c.Params("id")

	err = db.Where("id = ?",id).First(&posts).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "faild", "message":"post not found"})
	}

	if user.ID != posts.UserID {
		return c.Status(403).JSON(fiber.Map{"status": "faild", "message": "you don't have permission to delete this post"})
	}

	posts.IsActive = false
	err = db.Save(&posts).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "couldnt save"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "post deleted successfully"})
}

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




