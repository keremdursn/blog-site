package controllers

import (
	"app/database"
	"app/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	var user model.User
	c.BodyParser(&user)

	err := database.DB.Db.Create(&user).Error
	if err != nil {
		log.Println("user oluşturulamadi")
		return err
	}
	log.Println("user oluşturuldu")
	return nil
}

func LogIn(c *fiber.Ctx) error {

	var login model.Login
	// c.BodyParser(&login)

	// log.Println("**********", username)
	// log.Println("**********", password)
	var user model.User
	err := database.DB.Db.Where("username =? and password =? ", login.Username, login.Password).First(&user).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failed", "message": "user not found"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "login success", "data": user})
}

func tester() {
	log.Println("test yapildi")
}
