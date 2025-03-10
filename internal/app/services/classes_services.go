package services

import (
	"backend/internal/app/models"
	"backend/internal/app/repositories"
	"errors"
	"log"
)

type RegisteredList struct {
	categoly           string
	classIDs []int
}

type ClassesServise interface {
	RegisterUserClasses(classes *models.UserClasses) error
	ResponseUserClasses(filter *repositories.UserDetail) []models.Classes
	ResponseRegisteredClasses(userid string) []models.UserClasses
	CheckRegiseredClasses(userid string) (bool, []RegisteredList,error)
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

func (s *classesServise) ResponseUserClasses(filter *repositories.UserDetail) []models.Classes {
	classes, err := s.rep.GetUserClasses(filter)
	if err != nil {
		return nil
	}
	return classes
}

func (s *classesServise) ResponseRegisteredClasses(userid string) []models.UserClasses {
	classes, err := s.rep.GetRegisteredClasses(userid)
	if err != nil {
		return nil
	}
	return classes
}

func (s *classesServise) CheckRegiseredClasses(userid string) (bool, []RegisteredList,error) {
	categolyMap := map[string][]int{
		"IsIntroductoryList":{},
		"IsCoreList":{},
		"IsCommonList":{},
	}

	userClasses, err := s.rep.GetRegisteredClasses(userid)
	if err != nil {
		log.Println("GetRegisteredClasses in error :", err)
		return false, []RegisteredList{},err
	}

	classIDlist := make([]int,len(userClasses))
	for i,uc := range userClasses {
		classIDlist[i] = uc.ClassId
	}

	classes, err := s.rep.GetClassesByClassID(classIDlist)
	if err != nil {
		log.Println("GetClassesByClassID in error:", err)
		return false, []RegisteredList{},err
	}

	for _,class := range classes {
		switch {
		case class.IsIntroductory:
			categolyMap["IsIntroductoryList"] = append(categolyMap["IsIntroductoryList"],class.ClassId)
		case class.IsCore:
			categolyMap["IsCore"] = append(categolyMap["IsCore"],class.ClassId)
		case class.IsCommon:
			categolyMap["IsCommon"] = append(categolyMap["IsCommon"],class.ClassId)
		}
	}


	result,registeredList := varidataLists(categolyMap)
	if result {
		return result,registeredList,nil
	}
	return false,[]RegisteredList{},err
}



func varidataLists(categolyMap map[string][]int) (bool, []RegisteredList) {
	var registeredList []RegisteredList

	minRequirement := map[string]int {
		"IsIntroductory":2,
		"IsCore":        3,
		"IsCommon":      5,
	}


	for categoly,classIDlist := range categolyMap {
		if len(classIDlist) < minRequirement[categoly] {
			registeredList = append(registeredList,
				 RegisteredList{
					categoly: categoly,
					classIDs: classIDlist,
			})
		}
	}

	return len(registeredList) == 0,registeredList
}
