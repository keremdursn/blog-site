package middleware

import (
	"app/database"
	"app/model"

	"github.com/gofiber/fiber/v2"
)

func TokenControl(c *fiber.Ctx) (model.User, error) {
	db := database.DB.Db
	authorizationHeader := c.Get("Authorization")

	if authorizationHeader == "" || len(authorizationHeader) < 7 || authorizationHeader[:7] != "Bearer " {
		return model.User{}, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or missing token",
		})
	}
	token := authorizationHeader[7:]
	var user model.User
	session := new(model.Session)
	err := db.Where("token =?", token).First(&session).Error
	if err != nil {
		return model.User{}, c.Status(500).JSON(fiber.Map{"status": "error", "message": "you don't have session", "data": err})
	}

	userID := session.UserID
	err = db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return model.User{}, c.Status(500).JSON(fiber.Map{"status": "error", "message": "user not found!", "data": err})
	}

	return user, nil

}
