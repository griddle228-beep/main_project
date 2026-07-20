package controllers

import (
	"net/http"
	"strconv"
	"github.com/jackc/pgx/v5"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetAllFriends(ctx *gin.Context) {
	idparam := ctx.Param("id")
	userID, err := strconv.Atoi(idparam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id должен быть положительным числом"})
		return
	}
	friends, err := h.DbPool.GetAllUserFriends(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении друзей"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"friends": friends})
}
func (h *Handlers) GetAllMyFriends(ctx *gin.Context) {
	value, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return	
	}
	userID := value.(int)
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id должен быть положительным числом"})
		return
	}
	friends, err := h.DbPool.GetAllUserFriends(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении всех друзей"})
		return		
	}
	ctx.JSON(http.StatusOK, gin.H{"friends": friends})
}
func (h *Handlers) GetCountFriends(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id должен быть положительным числом"})
		return
	}
	countFriends, err := h.DbPool.GetCountFriends(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении колличества друзей"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"count_friends": countFriends})
}
func (h *Handlers) CheckFriendship(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userSecondID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return
	}
	if userSecondID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id должен быть положительным числом"})
		return
	}
	value, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат id"})
		return	
	}
	userID := value.(int)
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не авторизован"})
		return
	}
	friendship, err := h.DbPool.GetFriendship(userID, userSecondID)
	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusOK, gin.H{"message": "Данный пользователь не является вашим другом", "friendship": friendship})
		return	
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении дружеских отношений"})
		return		
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Данный пользователь является вашим другом", "friendship": friendship})
}
