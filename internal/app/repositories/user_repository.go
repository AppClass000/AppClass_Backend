package repositories

import (
	"backend/internal/app/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.Users, error)
	GetUserByUserID(userID string) (*models.Users, error)
	CreateUsers(user *models.Users) error
	RegisterUserName(users *models.Users, userName string) error
	RegisterUserDetail(userdetail *models.UserDetail) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByEmail(email string) (*models.Users, error) {
	var user models.Users
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println("error in getUserEmail:", err)
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUserID(userID string) (*models.Users, error) {
	var user models.Users
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUsers(user *models.Users) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("ユーザー作成エラー:%w", err)
	}
	return nil
}

func (r *userRepository) RegisterUserName(users *models.Users, userName string) error {
	if err := r.db.Model(users).Update("name", userName).Error; err != nil {
		return fmt.Errorf("ユーザー作成エラー:%w", err)
	}
	return nil
}

func (r *userRepository) RegisterUserDetail(userdetail *models.UserDetail) error {
	err := r.db.Create(userdetail).Error
	if err != nil {
		return fmt.Errorf("error in create userdetail:%w", err)
	}
	return nil
}
