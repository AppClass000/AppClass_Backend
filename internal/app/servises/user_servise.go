package servises

import (
	"backend/internal/app/models"
	"backend/internal/app/repositories"
	"backend/pkg/utils"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserServise interface {
	ResisterUser(User *models.Users) error
	HashUserPassword(User *models.Users) (*models.Users,error)
	ResponseUserIDJWT(email string,password string) (string ,error)
}

type userServise struct {
	rep repositories.UserRepository
}

func NewUserServise (rep *repositories.UserRepository) UserServise {
	return &userServise{rep: *rep}
}

func (s *userServise) ResisterUser(User *models.Users) error {
	if err := s.rep.Create(User);nil != err {
		return fmt.Errorf("ユーザー生成クエリエラー")
	}
	return nil
}

func (s *userServise) HashUserPassword(User *models.Users) (*models.Users,error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
    if err != nil {
	    return nil, fmt.Errorf("ハッシュ化失敗")
    }
    User.Password = string(hashed)
	return User ,nil
}


func (s *userServise) ResponseUserIDJWT(Email string,password string) (string,error) {
	users,err := s.rep.GetUserByEmail(Email);
	if err != nil {
		fmt.Errorf("missing Get User By Email %w",err)
	}
	if users.Password != password {
		fmt.Errorf("password is missing")
	}
	tokenstring,err := utils.GenerateJWT(users.UserId)
	if err != nil {
		fmt.Errorf("error in JWT generation")
	}
	return tokenstring,nil
}
