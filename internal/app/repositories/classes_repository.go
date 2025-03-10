package repositories

import (
	"backend/internal/app/models"
	"gorm.io/gorm"
	"fmt"
	"log"

)


type UserDetail struct {
	UserID    int
	Faculty    string 
	Department string 
	Course     string
	
}

type ClassesRepository interface {
	GetUserClasses(userdetail *UserDetail) ([]models.Classes,error)
	GetRegisteredClasses(userid string) ([]models.UserClasses, error)
	GetClassesByClassID(classIDList []int) ([]models.Classes, error)
	Create(userClass *models.UserClasses) error
}


type classesRepository struct {
    db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassesRepository {
	return &classesRepository{db: db}
}

func (r *classesRepository ) GetUserClasses(userdetail *UserDetail) ([]models.Classes, error) {
	var classes []models.Classes

	query := r.db.Select(
		"class_name", 
		"class_id",
		"is_mandatory",
		"is_core","is_",
		"is_introductory",
		"is_common",
		"genre",
		"semester",
		"instructor",
		"location",
		"schedule",
	)
	if userdetail.Faculty != "" {
		query = query.Where("faculty = ? or faculty = ?",userdetail.Faculty,"全学部")
	}
	err := query.Find(&classes).Error
	if err != nil {
		return nil, fmt.Errorf("授業情報を取得できませんでした")
	}

	return classes, nil
}

func (r *classesRepository) GetRegisteredClasses(userid string) ([]models.UserClasses, error) {
	var RegisteredClasses []models.UserClasses

	query := r.db.Select("class_name", "class_id","schedule")
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


func (r *classesRepository) GetClassesByClassID(classIDList []int) ([]models.Classes, error) {
	var RegisteredClasses []models.Classes
	if  len(classIDList) != 0 {
		return nil,fmt.Errorf("list of classID is 0")
	}

	query := r.db.Select("is_core","is_introductory","is_common","genre").Where("class_id IN",classIDList)

	log.Println("Executing SQL",query.Statement.SQL.String())
	
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
