package controllers

import (
	"log"
	"net/http"
	"semen_project/internal/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// user
func (h *Handlers) GetUserById(ctx *gin.Context) {
}
func (h *Handlers) SearchUsers(ctx *gin.Context) {
}
func (h *Handlers) UpdateUser(ctx *gin.Context) {
}
func (h *Handlers) UpdatePassword(ctx *gin.Context) {
}





















func (h *Handlers) CreateUser(ctx *gin.Context) {
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
		log.Printf("ОШИБКА в CreateUser: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": createdUser})
}

func (h *Handlers) GetAllUsers(ctx *gin.Context) {
	аllUsers, err := h.DbPool.GetAllUsers()
	if err != nil {
		log.Printf("ОШИБКА в GetAllUsers: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve users"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": аllUsers})
}
func (h *Handlers) Authentication(ctx *gin.Context) {
	var loginData struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	if err := ctx.BindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
		return
	}
	if strings.TrimSpace(loginData.UserName) == "" || strings.TrimSpace(loginData.Password) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "UserName и Password не могут быть пустыми"})
		return
	}

	user, err := h.DbPool.GetUserByUsername(loginData.UserName)
	if err != nil {
		log.Printf("ОШИБКА в GetUserByUsername: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	if user.Password != loginData.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учётные данные"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
func (h *Handlers) GetAllFriends(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	аllFriends, err := h.DbPool.GetAllFriends(userID)
	if err != nil {
		log.Printf("ОШИБКА в GetAllFriends: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve friends",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"friends": аllFriends})
}
func (h *Handlers) CreatePost(ctx *gin.Context) {
	var post models.Post
	if err := ctx.BindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if strings.TrimSpace(post.Content) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Content не может быть пустым"})
		return
	}
	if err := h.DbPool.CreatePost(post.UserID, post.Content); err != nil {
		log.Printf("ОШИБКА в CreatePost: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create post"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}
func (h *Handlers) GetAllPosts(ctx *gin.Context) {
	аllPosts, err := h.DbPool.GetAllPosts()
	if err != nil {
		log.Printf("ОШИБКА в GetAllPosts: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve posts",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Posts": аllPosts})
}
