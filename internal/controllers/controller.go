package controllers

import (
	"log"
	"net/http"
	"semen_project/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)
// user
func (h *Handlers) GetUserById(ctx *gin.Context) {
}
func (h *Handlers) SearchUsers(ctx *gin.Context) {
}
func (h *Handlers) GetUserByUsername(ctx *gin.Context) {
}
func (h *Handlers) GetUserByFirstName(ctx *gin.Context) {
}
func (h *Handlers) GetUserByLastName(ctx *gin.Context) {
}
func (h *Handlers) GetUserByFullName(ctx *gin.Context) {
}
// post
func (h *Handlers) CreatePost(ctx *gin.Context) {
}
func (h *Handlers) GetPostById(ctx *gin.Context) {
}
func (h *Handlers) GetFeed(ctx *gin.Context) {
}
func (h *Handlers) GetAllPosts(ctx *gin.Context) {
}
func (h *Handlers) GetAllUserPosts(ctx *gin.Context) {
}
func (h *Handlers) DeletePost(ctx *gin.Context) {
}
func (h *Handlers) UpdatePost(ctx *gin.Context) {
}
// like
func (h *Handlers) LikePost(ctx *gin.Context) {
}
func (h *Handlers) DeleteLike(ctx *gin.Context) {
}
func (h *Handlers) GetAllLikes(ctx *gin.Context) {
}
func (h *Handlers) CountLkes(ctx *gin.Context) {
}
// comment
func (h *Handlers) CreateComment(ctx *gin.Context) {
}
func (h *Handlers) DeleteComment(ctx *gin.Context) {
}
func (h *Handlers) UpdateComment(ctx *gin.Context) {
}
func (h *Handlers) GetallComments(ctx *gin.Context) {
}
func (h *Handlers) GetCountComments(ctx *gin.Context) {
}
// chat
func (h *Handlers) CreateChat(ctx *gin.Context) {
}
func (h *Handlers) SendMessage(ctx *gin.Context) {
}
func (h *Handlers) GetAllChats(ctx *gin.Context) {
}
func (h *Handlers) DeleteMessage(ctx *gin.Context) {
}
func (h *Handlers) DeleteChat(ctx *gin.Context) {
}
func (h *Handlers) GetAllMessages(ctx *gin.Context) {
}
func (h *Handlers) GetMarkRead(ctx *gin.Context) {
}
func (h *Handlers) GetCountNotReadMessage(ctx *gin.Context) {
}
// notification
func (h *Handlers) GetAllNotifications(ctx *gin.Context) {
}
func (h *Handlers) GetNotification(ctx *gin.Context) {
}
func (h *Handlers) CreateNotification(ctx *gin.Context) {
}
func (h *Handlers) DeleteNotification(ctx *gin.Context) {
}
// authentication
func (h *Handlers) Register(ctx *gin.Context) {
}
func (h *Handlers) Login(ctx *gin.Context) {
}
func (h *Handlers) RefreshToken(ctx *gin.Context) {
}
func (h *Handlers) Logout(ctx *gin.Context) {
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
