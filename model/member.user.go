package model

type User struct {
	UserId   uint   `gorm:"primarykey"`
	UserName string `json:"user_name"`
}

type Users struct {
	Users []User `json:"users"`
}
