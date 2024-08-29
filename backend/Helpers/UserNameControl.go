package helpers

import (
	"app/database"
	"app/model"
)

func UserNameControl(username string) error {
	var user model.User
	err := database.DB.Db.Where("username = ? ", username).First(&user).Error

	if err != nil {
		return err
	}
	return nil
}
