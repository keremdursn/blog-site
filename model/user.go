package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
