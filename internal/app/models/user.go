package models

import "time"

type Users struct {
    ID        uint      `gorm:"primaryKey;autoIncrement"` 
    UserID    string    `gorm:"column:user_id;unique"`
    Name      string    `gorm:"column:name;not null"`
    Email     string    `gorm:"column:email;unique;not null"`
    Password  string    `gorm:"column:password;not null"`
    CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}


type UserDetail struct {
	ID        uint   `gorm:"primarykey"`
	UserID    string `gorm:"primarykey"`
	Faculty    string `gorm:"not null"`
	Department string `gorm:"not null"`
	Course     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ReqUsers struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestUserID struct {
	ReqUserID string
}