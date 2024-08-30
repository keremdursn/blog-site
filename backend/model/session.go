package model

type Session struct {
	UserID   uint `gorm:"primaryKey;autoIncrement`
	Token    string
	IsActive bool `gorm:"default:true"`
}
