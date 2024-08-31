package middleware

import (
	"app/database"
	"app/model"

	"github.com/gofiber/fiber/v2"
)

func TokenControl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db := database.DB.Db
		authorizationHeader := c.Get("Authorization")
		if authorizationHeader == "" || len(authorizationHeader) < 7 || authorizationHeader[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or missing token",
			})
		}
		token := authorizationHeader[7:]

		session := new(model.Session)
		err := db.Where("token =?", token).First(&session).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "you don't have session", "data": err})
		}

		user := new(model.User)
		err = db.Where("id=? ", session.UserID).First(&user).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "user not found", "data": err})
		}
		c.Locals("user", user)

		return c.Next()
	}

}
