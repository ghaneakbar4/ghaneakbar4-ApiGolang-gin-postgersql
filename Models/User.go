package models

//ساختار یوزر
type User struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Age uint32 `json:"age"`
	Address string `json:"address"`
}

type UserUri struct {
	ID uint `uri:"id" binding:"required,number"`
}