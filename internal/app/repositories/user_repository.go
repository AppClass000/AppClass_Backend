package repositories

ty

func FindEmailUser(db *gorm.DB, email string) (*Users, error) {
	var user Users
	if err := db.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(db *gorm.DB, user *Users) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		log.Println("全ての項目が入力されていません")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("ハッシュ化失敗")
	}
	user.Password = string(hashed)

	if err := db.Create(user).Error; err != nil {
		return fmt.Errorf("ユーザー作成エラー:%w", err)
	}
	return nil
}

func GetUserDetail(db *gorm.DB,userID string) (*UserDetail,error)  {
	var userDetail UserDetail
	if err := db.Select("user_id",userID).First(&userDetail).Error;err != nil {
		return nil,err 
	}
	return &userDetail,nil 
}