package controllers

import (
	"net/http"
	"strconv"
	"github.com/jackc/pgx/v5"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) LikePost(ctx *gin.Context) {
	idparam := ctx.Param("id")
	postID, err := strconv.Atoi(idparam)
	if err != nil || postID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id поста"})
		return
	}

	value, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении id из токена"})
		return
	}

	userID := value.(int)
	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id пользователя"})
		return
	}

	_, err = h.DbPool.GetPostById(postID)
	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении поста"})
		return
	}

	liked, err := h.DbPool.GetLikeStatus(postID, userID)
	if err != nil && err != pgx.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке лайка"})
		return
	}

	if liked {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Вы уже поставили лайк"})
		return
	}

	err = h.DbPool.CreateLike(postID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании лайка"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Лайк поставлен"})
}
func (h *Handlers) DeleteLike(ctx *gin.Context) {
	idparam := ctx.Param("id")
	postID, err := strconv.Atoi(idparam)
	if err != nil || postID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id поста"})
		return
	}

	value, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении id из токена"})
		return
	}

	userID := value.(int)

	_, err = h.DbPool.GetPostById(postID)
	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	_, err = h.DbPool.GetLikeStatus(postID, userID)
	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Лайк не найден"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке лайка"})
		return
	}

	err = h.DbPool.DeleteLike(postID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении лайка"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Лайк удален"})
}
func (h *Handlers) GetAllPostLikes(ctx *gin.Context) {
	idparam := ctx.Param("id")
	postID, err := strconv.Atoi(idparam)
	if err != nil || postID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id поста"})
		return
	}

	_, err = h.DbPool.GetPostById(postID)
	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	if err != nil {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении поста"})
	return
	}

	likes, err := h.DbPool.GetAllPostLikes(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении лайков"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
func (h *Handlers) GetCountLikes(ctx *gin.Context) {
	idparam := ctx.Param("id")
	postID, err := strconv.Atoi(idparam)
	if err != nil || postID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id поста"})
		return
	}

	_, err = h.DbPool.GetPostById(postID)
	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}
	if err != nil {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении поста"})
	return
	}

	count, err := h.DbPool.GetCountLikes(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении количества лайков"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"count": count})
}
func (h *Handlers) GetAllUserLikes(ctx *gin.Context) {
	idparam := ctx.Param("id")
	userID, err := strconv.Atoi(idparam)
	if err != nil || userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id пользователя"})
		return
	}

	_, err = h.DbPool.GetUserById(userID)
	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}
	if err != nil {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении поста"})
	return
	}

	likes, err := h.DbPool.GetAllUserLikes(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении лайков пользователя"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
func (h *Handlers) GetLikeStatus(ctx *gin.Context) {
	idparam := ctx.Param("id")

	postID, err := strconv.Atoi(idparam)
	if err != nil || postID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id поста"})
		return
	}

	value, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка при получении id из токена"})
		return
	}

	userID := value.(int)

	if userID <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный id пользователя"})
		return
	}

	_, err = h.DbPool.GetPostById(postID)
	if err == pgx.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Пост не найден"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении поста"})
		return
	}

	liked, err := h.DbPool.GetLikeStatus(postID, userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при проверке лайка"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"liked": liked,
	})
}