package repositories

import (
	"backend/internal/app/models"
	"gorm.io/gorm"
	"backend/internal/db"
	"fmt"
	"log"

)

type ClassesRepository interface {
	GetClasses(filter *models.UserDetail) ([]models.Classes,error)
	GetUserClasses(userid string) ([]models.UserClasses, error)
	Create(userClass *models.UserClasses) error
}


type classesRepository struct {
    db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassesRepository {
	return &classesRepository{db}
}

func (r *classesRepository ) GetClasses(filter *models.UserDetail) ([]models.Classes, error) {
	var classes []models.Classes

	query := db.DB.Select("class_name", "class_id", "is_mandatory", "instructor", "location", "schedule")
	if filter != nil {
		query = query.Where("faculty = ?", filter.Faculty).Or("faculty = ?", "全学部")
	}
	err := query.Find(&classes).Error
	if err != nil {
		return nil, fmt.Errorf("授業情報を取得できませんでした")
	}

	return classes, nil
}

func (r *classesRepository)GetUserClasses(userid string) ([]models.UserClasses, error) {
	var userClasses []models.UserClasses

	query := db.DB.Select("class_name", "class_id", "is_mandatory", "instructor", "location", "schedule")
	if userid != "" {
		query = query.Where("user_id = ?", userid)
		log.Println(query)
	}
	err := query.Find(&userClasses).Error
	if err != nil {
		return nil, fmt.Errorf("ユーザー授業情報を取得できませんでした")
	}

	return userClasses, nil

}

func (r *classesRepository) Create(userClass *models.UserClasses) error {

	err := db.DB.Create(userClass).Error
	if err != nil {
		log.Println("userClassのレコード作成失敗", err)
		return err
	}

	return nil
}
