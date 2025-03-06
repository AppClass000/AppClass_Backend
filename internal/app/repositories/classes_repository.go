package repositories

import (
	"backend/internal/app/models"
	"gorm.io/gorm"
	"fmt"
	"log"

)

type ClassesRepository interface {
	GetUserClasses(filter *models.UserDetail) ([]models.Classes,error)
	GetRegisteredClasses(userid string) ([]models.ClassesDetail, error)
	Create(userClass *models.UserClasses) error
}


type classesRepository struct {
    db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassesRepository {
	return &classesRepository{db: db}
}

func (r *classesRepository ) GetUserClasses(filter *models.UserDetail) ([]models.Classes, error) {
	var classes []models.Classes

	query := r.db.Select("class_name", "class_id", "is_mandatory", "instructor", "location", "schedule")
	if filter != nil {
		query = query.Where("faculty = ? or faculty = ?",filter.Faculty,"全学部")
	}
	err := query.Find(&classes).Error
	if err != nil {
		return nil, fmt.Errorf("授業情報を取得できませんでした")
	}

	return classes, nil
}

func (r *classesRepository) GetRegisteredClasses(userid string) ([]models.ClassesDetail, error) {
	var RegisteredClasses []models.ClassesDetail

	query := r.db.Select("class_name", "class_id", "is_mandatory", "is_core", "is_introductory", "is_common")
	if userid != "" {
		query = query.Where("user_id = ?", userid)
		log.Println("Executing SQL",query.Statement.SQL.String())
	}
	err := query.Find(&RegisteredClasses).Error
	if err != nil {
		return nil, fmt.Errorf("ユーザー授業情報を取得できませんでした")
	}

	return RegisteredClasses, nil

}

func (r *classesRepository) Create(userClass *models.UserClasses) error {

	err := r.db.Create(userClass).Error
	if err != nil {
		log.Println("userClassのレコード作成失敗", err)
		return err
	}

	return nil
}
