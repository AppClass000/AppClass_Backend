package handlers

import (
	"backend/internal/app/models"
	"backend/internal/app/services"
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
	var filter models.UserDetail

	err := c.ShouldBindJSON(&filter)
	if err != nil {
		log.Println("missing Bind UserDetail:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "missing Bind UserDetail",
		})
		return
	}
	classes := h.serv.ResponseUserClasses(&filter)

	c.JSON(http.StatusOK, gin.H{
		"message":     "success",
		"userClasses": classes,
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
