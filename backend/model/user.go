package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.New()
	return
}

type UpdateUser struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
}

type ChangePassword struct {
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
