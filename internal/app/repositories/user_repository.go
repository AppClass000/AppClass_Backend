package repositories

import (
	"backend/internal/app/models"
	"backend/internal/database"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.Users, error)
	Create(user *models.Users) error 
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db:db}
}

func (r *userRepository) GetUserByEmail(email string) (*models.Users, error) {
	var user models.Users
	if err := r.db.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *models.Users) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("ユーザー作成エラー:%w", err)
	}
	return nil
}

func  GetUserDetail(userID string) (*models.UserDetail,error)  {
	var userDetail models.UserDetail
	if err := database.DB.Select("user_id",userID).First(&userDetail).Error;err != nil {
		return nil,err 
	}
	return &userDetail,nil 
}