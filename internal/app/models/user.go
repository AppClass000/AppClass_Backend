package models

import "time"

type Users struct {
	Id       uint   `gorm:"primarykey"`
	UserId   string `gorm:"unique;column:UserId"`
	Name     string `gorm:"not null" json:"userName"`
	Email    string `gorm:"unique;not null" json:"userEmail"`
	Password string `gorm:"not null" json:"userPassword"`
}

type UserDetail struct {
	Id         uint   `gorm:"primarykey"`
	UserId     string `gorm:"primarykey"`
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
	ReqUserId string
}