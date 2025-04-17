package handlers

import (
	"backend/internal/app/models"
	"backend/internal/app/repositories"
	"backend/internal/app/services"
	"backend/pkg/utils"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ClassesHandler struct {
	serv services.ClassesServise
}

func NewClassesHandler(serv services.ClassesServise) ClassesHandler {
	return ClassesHandler{serv: serv}
}

func (h *ClassesHandler) ViewUserClasses(c *gin.Context) {
	var filter repositories.UserDetail

	err := c.ShouldBindJSON(&filter)
	if err != nil {
		log.Println("missing Bind UserDetail:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "missing Bind UserDetail",
		})
		return
	}
	log.Println("this is filter:", filter)

	classes := h.serv.ResponseUserClasses(&filter)

	c.JSON(http.StatusOK, gin.H{
		"message":     "success",
		"userClasses": classes,
	})
}

func (h *ClassesHandler) ViewUserClassesByUserID(c *gin.Context) {
	value, exist := c.Get("userID")
	if !exist {
		log.Println("UserIDがありません:classes_handler")
	}
	userId, ok := value.(string)
	if !ok {
		log.Println("Invalid UserID Type")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid userid",
		})
		return
	}
	userDetail, err := utils.GetUserDetail(userId)
	if err != nil {
		log.Println("missing get userDetail by userid:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "missing get userDetail by userid",
		})
		return
	}
	userClasses := h.serv.ResponseUserClasses(userDetail)

	c.JSON(http.StatusOK, gin.H{
		"message":     "success",
		"userclasses": userClasses,
	})

}

func (h *ClassesHandler) RegisterClass(c *gin.Context) {
	var userClass models.UserClasses
	value, exist := c.Get("userID")
	if !exist {
		log.Println("UserIDがありません:classes_handler")
	}
	userId, ok := value.(string)
	if !ok {
		log.Println("Invalid UserID Type")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid userid",
		})
		return
	}

	err := c.ShouldBindJSON(&userClass)
	if err != nil {
		log.Println("missing Bind userClass:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "missing Bind userClass",
		})
		return
	}
	userClass.UserID = userId

	err = h.serv.RegisterUserClasses(&userClass)
	if err != nil {
		log.Println("occur error in registerUserClasses")
		c.JSON(204, gin.H{
			"error": "This class is registered ",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "registered success",
	})
}

func (h *ClassesHandler) DeleteRegisteredClass(c *gin.Context) {
	var userClass models.UserClasses

	err := c.ShouldBindJSON(&userClass)
	if err != nil {
		log.Println("missing Bind userClass:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "missing Bind userClass",
		})
		return
	}
	log.Println(userClass)
	err = h.serv.DeleteRegisteredClasses(&userClass)
	if err != nil {
		log.Println("error in  deleteRegisteredClass:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error in  deleteRegisteredClass:",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success delete  userClass",
	})
}

func (h *ClassesHandler) ViewUserSchedule(c *gin.Context) {
	value, exist := c.Get("userID")
	if !exist {
		log.Println("UserIDがありません:classes_handler")
	}
	userId, ok := value.(string)
	if !ok {
		log.Println("Invalid UserID Type")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid userid",
		})
		return
	}

	registeredClasses := h.serv.ResponseRegisteredClasses(userId)
	c.JSON(http.StatusOK, gin.H{
		"message":           "success",
		"registeredClasses": registeredClasses,
	})

}

func (h *ClassesHandler) CheckToolAPI(c *gin.Context) {
	value, exist := c.Get("userID")
	if !exist {
		log.Println("UserIDがありません:classes_handler")
	}
	userId, ok := value.(string)
	if !ok {
		log.Println("Invalid UserID Type")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid userid",
		})
		return
	}

	result, registeredList, err := h.serv.CheckRegiseredClasses(userId)
	if err != nil {
		log.Println("error in CheckRegisteredClasses", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "success",
		"checktoolresult": result,
		"registeredlist":  registeredList,
	})

}
