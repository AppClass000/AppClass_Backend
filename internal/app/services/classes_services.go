package services

import (
	"backend/internal/app/models"
	"backend/internal/app/repositories"
	"log"
	"errors"
)

type ClassesServise interface {
	RegisterUserClasses(classes *models.UserClasses) error
	ResponseUserClasses(filter *models.UserDetail) []models.Classes
	ResponseRegisteredClasses(userid string) []models.ClassesDetail
}

type classesServise struct {
	rep repositories.ClassesRepository
}

func NewClassesServise(rep repositories.ClassesRepository) ClassesServise {
	return &classesServise{rep: rep}
}

func (s *classesServise) RegisterUserClasses(classes *models.UserClasses) error {
	if classes.UserId == "" {
		log.Println("userID does not exist")
		return errors.New("UserID is required")
	}
	err := s.rep.Create(classes)
	if err != nil {
		return err
	}
	log.Println("success registered classes for users")
	return nil
}

func (s *classesServise) ResponseUserClasses(filter *models.UserDetail) []models.Classes {
	classes, err := s.rep.GetUserClasses(filter)
	if err != nil {
		return nil
	}
	return classes
}

func (s *classesServise) ResponseRegisteredClasses(userid string) []models.ClassesDetail {
	classes, err := s.rep.GetRegisteredClasses(userid)
	if err != nil {
		return nil
	}
	return classes
}

// func CheckUserResitered(c *gin.Context) {
// 	var IsCoreList []int
// 	var IsIntroductoryList []int
// 	var IsCommonList []int

//  	UserID, ok := c.Get("userID")
//  	if !ok {
//  		log.Panic("userID don't exist")
//  	}
// 	UserId, ok := UserID.(string)
//  	if !ok {
//  		log.Println("mistake convert UserID to string")
//  	}

// // 	userClasses, err := models.GetUserClasses(UserId)
// // 	if err != nil {
// // 		log.Println("Getting userCLasses error")
// // 	}

// // 	for i :=0; i< len(userClasses); i++ {
// // 		if userClasses[i].IsCore == 1 {
// // 			IsCoreList := append(IsCoreList,userClasses[i].IsCore)
// // 		}
// // 	    if userClasses[i].IsIntroductory == 1 {
// // 			IsIntroductoryList := append(IsIntroductoryList,userClasses[i].IsIntroductory)
// // 		}
// // 	    if userClasses[i].IsCommon == 1 {
// // 	    	IsCommonList := append(IsCommonList,userClasses[i].IsCommon)
// // 	    }
// // 	}

// // 	userDetail,err := models.GetUserDetail(database.DB,UserId)
// // 	if err != nil {
// // 		log.Println("error of GetUserDetail:",err)
// // 	}
// // 	if userDetail.Faculty == "工学部" {

// // 	}

// // }
// // //
