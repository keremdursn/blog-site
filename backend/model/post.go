package model

type Post struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      uint   `json:"user_id"`
	User        User   `gorm:"foreign:UserID"`
}
