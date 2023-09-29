package model

import "time"

type Password struct {
	PasswordId uint      `gorm:"primarykey"`
	UserId     uint      `json:"user_id"`
	Password   string    `json:"password"`
	Updated    time.Time `json:"updated"`
}
