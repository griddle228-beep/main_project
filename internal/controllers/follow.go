package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func (h *Handlers) FollowUser(ctx *gin.Context) {
	value, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID не найден в контексте"})
		return
	}
	userID := value.(int)
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	idparam := ctx.Param("id")
	followingUserID, err := strconv.Atoi(idparam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	if followingUserID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	if userID == followingUserID {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": "нельзя подписаться на самого себя"})
    return
	}
	_ , err = h.DbPool.GetUserById(followingUserID)
	if err == pgx.ErrNoRows {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": "пользователя не существует"})
    return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке пользователя"})
		return
	}
	friendship, err := h.DbPool.GetFollowStatus(followingUserID, userID)
	if err != pgx.ErrNoRows && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при проверке подписки"})
		return
	}
	if friendship {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "вы уже подписаны на данного пользователя"})
		return
	}
	checkFollowingToUserStatus, err := h.DbPool.GetFollowStatus(userID, followingUserID)
	if err != pgx.ErrNoRows && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при проверке статуса подписки"})
		return
	}
	if checkFollowingToUserStatus {
		err = h.DbPool.UnFollowUser(followingUserID, userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при удалении подписки пользователя"})
			return
		}
		err = h.DbPool.CreateFriendship(userID, followingUserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при создании дружбы"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "вы успешно добавили пользователя в друзья"})
		return
	}
	err = h.DbPool.FollowUser(userID, followingUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при создании подписки"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Вы успешно подписались на пользователя"})
}
func (h *Handlers) UnFollowUser(ctx *gin.Context) {
	value, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userID не найден в контексте"})
		return
	}
	userID := value.(int)
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	
	idparam := ctx.Param("id")
	unFollowingUserID,err := strconv.Atoi(idparam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	if unFollowingUserID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	if userID == unFollowingUserID {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "вы не можете отписаться/подписаться на себя"})
		return
	}
	_, err = h.DbPool.GetUserById(unFollowingUserID)
	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Пользователя с таким id не существует"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении пользователя"})
		return
	}
	_, err = h.DbPool.GetFollowStatus(unFollowingUserID, userID)
	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Вы не подписаны на данного пользователя"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке подписки"})
		return
	}
	statusFollowingToUser, err := h.DbPool.GetFriendship(userID, unFollowingUserID)
	if err != nil && err != pgx.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении статуса подписки"})
		return
	}
	if statusFollowingToUser {
		err = h.DbPool.DeleteFriend(userID, unFollowingUserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении дружбы"})
			return
		}

		err = h.DbPool.FollowUser(unFollowingUserID, userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении пользователя в подписки"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Вы успешно отписались от пользователя"})
		return
	}
	err = h.DbPool.UnFollowUser(userID, unFollowingUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении подписки"})
		return	
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Вы успешно отписались от пользователя"})
}
func (h *Handlers) GetAllFollowers(ctx *gin.Context) {
	idparam := ctx.Param("id")
	userID, err := strconv.Atoi(idparam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	followers, err := h.DbPool.GetAllUserFollowers(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении всех подписчиков"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"followers": followers})
}
func (h *Handlers) GetAllFollowing(ctx *gin.Context) {
	idparam := ctx.Param("id")
	userID, err := strconv.Atoi(idparam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	following, err := h.DbPool.GetAllUserFollowing(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении всех подписок"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"followings": following})
}
func (h *Handlers) GetCountFollowers(ctx *gin.Context) {
	idparam := ctx.Param("id")
	userID, err := strconv.Atoi(idparam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	count, err := h.DbPool.GetCountFollowers(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении колличества подписчиков"})
		return		
	}
	ctx.JSON(http.StatusOK, gin.H{"followers_count": count})
}
func (h *Handlers) GetCountFollowing(ctx *gin.Context) {
	idparam := ctx.Param("id")
	userID, err := strconv.Atoi(idparam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	count, err := h.DbPool.GetCountFollowing(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении колличества подписок"})
		return		
	}
	ctx.JSON(http.StatusOK, gin.H{"following_count": count})
}
func (h *Handlers) CheckFollowStatus(ctx *gin.Context) {
	idparam := ctx.Param("id")
	secondUserID, err := strconv.Atoi(idparam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	if secondUserID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	value, exists := ctx.Get("userID")	
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении id из токена"})
		return		
	}
	userID := value.(int)
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "некоректный id"})
		return
	}
	if secondUserID == userID {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Вы не можете получить статус подписки с самим собой"})
		return
	}
	_, err = h.DbPool.GetUserById(secondUserID)
	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Пользователя с таким id не существует"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении пользователя"})
		return
	}
	heFollowsMe, err := h.DbPool.GetFollowStatus(userID, secondUserID)
	if err != nil && err != pgx.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении статуса подписки пользователя"})
		return
	}
	iFollow, err := h.DbPool.GetFollowStatus(secondUserID,userID)
	if err != nil && err != pgx.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении статуса подписки пользователя"})
		return
	}
	var status string
	if iFollow && heFollowsMe {
		status = "friend"
	} else if iFollow {
		status = "following"
	} else if heFollowsMe {
		status = "follower"
	} else {
		status = "none"
	}
	ctx.JSON(http.StatusOK, gin.H{"status": status})
}