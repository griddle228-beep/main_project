package controllers

import (
	"net/http"
	"strings"
	"log"
	"semen_project/models"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": gin.H{
			"method": http.MethodGet,
			"code": http.StatusOK,
			"message": "qq",
		},
	})
}
func Answer(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": gin.H{
			"method": http.MethodGet,
			"code": http.StatusOK,
			"message": "👋👋👋",
		},
	})
}
func (h *Handlers) Create(ctx *gin.Context) {
	var err error
	var user models.User
	if err = ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if strings.TrimSpace(user.UserName) == "" || strings.TrimSpace(user.Password) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "UserName и Password не могут быть пустыми"})
		return
	}
	createdUser, err := h.DbPool.CreateUser(user)
if err != nil {
    // Добавьте подробное логирование ошибки
    log.Printf("ОШИБКА в CreateUser: %v", err)
    log.Printf("Данные пользователя: %+v", user)
    
    // Верните более информативную ошибку
    ctx.JSON(http.StatusInternalServerError, gin.H{
        "error": "Failed to create user",
        "details": err.Error(),  // Добавьте детали ошибки
    })
    return
}
	ctx.JSON(http.StatusOK, gin.H{"user": createdUser})
}