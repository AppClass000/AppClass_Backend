package models

type Classes struct {
	ID             int
	ClassName      string `gorm:"column:class_name"`
	ClassID        int    `gorm:"column:class_id"`
	Faculty        string `gorm:"not null"`
	Department     string
	Course         string
	IsMandatory    int    `gorm:"column:is_mandatory"`
	IsCore         bool   `gorm:"column:is_core"`
	IsIntroductory bool   `gorm:"column:is_introductory"`
	IsCommon       bool   `gorm:"column:is_common"`
	Instructor     string `gorm:"column:instructor"`
	Location       string `gorm:"column:location"`
	Schedule       string `gorm:"column:schedule"`
	Genre          string `gorm:"type:varchar(50)"`
	Semester       string `gorm:"type:varchar(2)"`
}

type UserClasses struct {
	Id         int
	UserID     string
	ClassName  string `gorm:"column:class_name"`
	ClassID    int    `gorm:"column:class_id"`
	Schedule   string `gorm:"column:schedule"`
	Instructor string `gorm:"column:instructor"`
	Location   string `gorm:"column:location"`
}

type ClassesDetail struct {
	ID             int    `gorm:"primaryKey"`
	ClassID        int    `gorm:"not null"`
	ClassName      string `gorm:"type:varchar(200)"`
	IsCore         bool   `gorm:"not null"`
	IsIntroductory bool   `gorm:"not null"`
	IsCommon       bool   `gorm:"not null"`
	Genre          string `gorm:"type:varchar(50)"`
	Semester       string `gorm:"type:varchar(2)"`
}
