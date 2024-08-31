package controllers

import (
	"app/database"
	"app/helpers"
	"app/middleware"
	"app/model"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User
	c.BodyParser(&user)
	// err := helpers.UserNameControl(user.Username)
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"status": "failed", "message": "username already taken."})
	// }

	user.Password = helpers.HashPass(user.Password)
	err := db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": " server error"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "signup success", "data": user})
}

func LogIn(c *fiber.Ctx) error {
	db := database.DB.Db
	login := new(model.Login)

	c.BodyParser(&login)

	login.Password = helpers.HashPass(login.Password)

	user := new(model.User)
	err := db.Where("username =? and password =? ", login.Username, login.Password).First(&user).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failed", "message": "user not found"})
	}

	token, err := middleware.GenerateJWT(login.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Token oluşturulamadı")
	}

	session := new(model.Session)
	session.UserID = user.ID
	session.Token = token
	db.Create(&session)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "login success", "data": user, "token": token})
}

func LogOut(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error"})
	}
	db := database.DB.Db
	session := new(model.Session)
	err = db.Where("user_id = ? ", user.ID).First(&session).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error"})
	}
	session.IsActive = false

	err = database.DB.Db.Raw("DELETE FROM sessions WHERE user_id= ?", user.ID).Scan(&session).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "logout success"})
}

func UpdateUser(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server errorrrrr"})
	}

	db := database.DB.Db
	// var user model.User  token kontrolü gelince bunu kaldırdık

	//İstekten kullanıcı ID sini al
	// id := c.Params("id")

	// err := db.Where("id = ?", id).First(&user).Error
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": "user not found"})
	// }

	// updateUserData := new(model.UpdateUser)
	// c.BodyParser(&updateUserData)
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{"status": "error", "message": "json bodyparse edilemedi.", "data": err})
	// }

	// user.Username = updateUserData.Username
	// user.Name = updateUserData.Name
	// user.Surname = updateUserData.Surname

	name := c.FormValue("name")
	surname := c.FormValue("surname")
	username := c.FormValue("username")

	if len(name) != 0 {
		user.Name = name
	}
	if len(surname) != 0 {
	    user.Surname = surname
	}
	if len(username) != 0 {
		user.Username = username
	}

	err = db.Save(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error!"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user data has been successfully updated "})

}

func ChangePassword(c *fiber.Ctx) error {
	// Kullanıcıyı token ile doğrula
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
	}

	// Şifre değişiklik isteğini parse et
	changePassword := new(model.ChangePassword)
	if err := c.BodyParser(&changePassword); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid request"})
	}

	// Eski şifreyi doğrula
	hashedOldPassword := helpers.HashPass(changePassword.OldPassword)
	if hashedOldPassword != user.Password {
		return c.Status(401).JSON(fiber.Map{"status": "faild", "message": "Old password is incorrect"})
	}

	// Yeni şifrelerin uyumunu kontrol et
	if changePassword.NewPassword1 != changePassword.NewPassword2 {
		return c.Status(401).JSON(fiber.Map{"status": "faild", "message": "New passwords do not match"})
	}

	// Şifrenin karmaşıklığını kontrol et (opsiyonel)
	if len(changePassword.NewPassword1) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Password must be at least 8 characters long"})
	}

	// Yeni şifreyi hashle
	hashedNewPassword := helpers.HashPass(changePassword.NewPassword1)

	// Şifreyi güncelle
	user.Password = hashedNewPassword
	db := database.DB.Db
	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to update password"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Password changed successfully"})
}

func DeleteAccount(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
	}

	db := database.DB.Db
	// var user model.User  token kontrolü gelince bunu kaldırdık

	// id := c.Params("id")
	// err = db.Where("id = ?", id).First(&user).Error
	// if err != nil {
	// 	return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": "user not found"})
	// }
 
	err = db.Delete(&user).Error
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
	_, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
	}
	db := database.DB.Db

	id := c.Params("id")
	var user model.User

	err = db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Status": "Error", "Message": "user not found"})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}
