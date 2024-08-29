package controllers

import (
	"app/database"
	"app/helpers"
	"app/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//password change

func SignUp(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User
	c.BodyParser(&user)
	err := helpers.UserNameControl(user.Username)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failed", "message": "username already taken."})
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": " server error"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "signup success", "data": user})
}

func LogIn(c *fiber.Ctx) error {
	db := database.DB.Db
	var login model.Login
	// c.BodyParser(&login)

	// log.Println("**********", username)
	// log.Println("**********", password)
	var user model.User
	err := db.Where("username =? and password =? ", login.Username, login.Password).First(&user).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failed", "message": "user not found"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "login success", "data": user})
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User

	//İstekten kullanıcı ID sini al
	id := c.Params("id")

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "user not found", "data": nil})
	}

	updateUserData := new(model.UpdateUser)
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "json bodyparse edilemedi.", "data": err})
	}

	user.Username = updateUserData.Username
	err = db.Updates(user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user data has been successfully updated "})

}

// func ChangePassword(c *fiber.Ctx) error {

// }

func DeleteAccount(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User

	id := c.Params("id")
	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data:": nil})
	}

	err := db.Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "failed to delete user", "data:": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "account has been successfully deleted. "})
}

func GetAllUser(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}

func GetUserByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User

	id := c.Params("id")

	db.Find(&user, "id =?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data:": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}
