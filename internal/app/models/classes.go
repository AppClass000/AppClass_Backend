package models


type Classes struct {
	Id          uint
	ClassName   string `gorm:"column:class_name"`
	ClassId     int    `gorm:"column:class_id"`
	IsMandatory int    `gorm:"column:is_mandatory"`
	Instructor  string `gorm:"column:instructor"`
	Location    string `gorm:"column:location"`
	Schedule    string `gorm:"column:schedule"`
}

type UserClasses struct {
	Id          uint
	UserId      string
	ClassName   string `gorm:"column:class_name"`
	ClassId     int    `gorm:"column:class_id"`
	IsMandatory int    `gorm:"column:is_mandatory"`
	IsCore         int    `gorm:"column:is_core"`
	IsIntroductory int    `gorm:"column:is_introductory"`
	IsCommon       int    `gorm:"column:is_common"`
}
