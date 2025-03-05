package servises

type ClassesServise interface {

}

type classesServise struct {
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