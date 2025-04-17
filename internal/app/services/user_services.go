package services

import (
	"backend/internal/app/models"
	"backend/internal/app/repositories"
	"backend/pkg/auth"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	*models.Users
}

type Task struct {
	UserID string
}
type Responses struct {
	User *models.Users
	Err  error
}

type UserServise interface {
	ResisterUser(User *models.Users) error
	HashUserPassword(User *models.Users) (*models.Users, error)
	ResponseUserIDJWT(email string, password string) (string, error)
	ResisterUserName(users *models.Users, userName string) error
	GetUserByUserID(userID string) (*models.Users, error)
	RegisterUserDetail(userdetail *models.UserDetail, userID string) error
	ResponseSignUpJTW(userID string) (string, error)
}

type userServise struct {
	rep repositories.UserRepository
}

func NewUserServise(rep *repositories.UserRepository) UserServise {
	return &userServise{rep: *rep}
}

func (s *userServise) ResisterUser(User *models.Users) error {
	if err := s.rep.CreateUsers(User); nil != err {
		return fmt.Errorf("ユーザー生成クエリエラー: %v", err)
	}
	return nil
}

func (s *userServise) RegisterUserDetail(userdetail *models.UserDetail, userID string) error {
	if userID == "" {
		return fmt.Errorf("userID is not exist")
	}

	userdetail.UserID = userID
	err := s.rep.RegisterUserDetail(userdetail)
	if err != nil {
		return fmt.Errorf("missing register userdetail: %v", err)
	}
	return nil
}

func (s *userServise) GetUserByUserID(userID string) (*models.Users, error) {
	user, err := s.rep.GetUserByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("unnable fetch user detail : %v", err)
	}

	return user, nil
}

func (s *userServise) ResisterUserName(users *models.Users, userName string) error {
	if userName == "" {
		return fmt.Errorf("not exist userName")
	}
	if err := s.rep.RegisterUserName(users, userName); nil != err {
		return fmt.Errorf("error in create userProfile")
	}
	return nil
}

func (s *userServise) HashUserPassword(User *models.Users) (*models.Users, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("ハッシュ化失敗")
	}
	User.Password = string(hashed)
	return User, nil
}

func (s *userServise) ResponseUserIDJWT(Email string, password string) (string, error) {
	if Email == "" {
		log.Println("email is  not exist")
	}
	users, err := s.rep.GetUserByEmail(Email)
	if err != nil {
		fmt.Errorf("missing Get User By Email %w", err)
	}
	if users.Password != password {
		fmt.Errorf("password is missing")
	}
	tokenstring, err := auth.GenerateJWT(users.UserID)
	if err != nil {
		fmt.Errorf("error in JWT generation")
	}
	return tokenstring, nil
}

func (s *userServise) ResponseSignUpJTW(userID string) (string, error) {
	tokenstring, err := auth.GenerateJWT(userID)
	if err != nil {
		fmt.Errorf("error in JWT generation")
	}
	return tokenstring, nil
}
