package controllers

import (
	"log"
	"net/http"
	"semen_project/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create user",
			"details": err.Error(), // Добавьте детали ошибки
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": createdUser})
}

func (h *Handlers) GetAllUsers(ctx *gin.Context) {
	аllUsers, err := h.DbPool.GetAllUsers()
	if err != nil {
		log.Printf("ОШИБКА в GetAllUsers: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve users",
			"details": err.Error(),
		})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve user",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
func (h *Handlers) GetAllFriends(ctx *gin.Context) {
	аllFriends, err := h.DbPool.GetAllFriends()
	if err != nil {
		log.Printf("ОШИБКА в GetAllFriends: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve friends",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Friends": аllFriends})
}
func (h *Handlers) CreatePost(ctx *gin.Context) {
	var post models.Post
	if err := ctx.BindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := h.DbPool.CreatePost(post.UserID, post.Content); err != nil {
		log.Printf("ОШИБКА в CreatePost: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create post",
			"details": err.Error(),
		})
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
func (h *Handlers) GetAllComments(ctx *gin.Context) {
	аllComments, err := h.DbPool.GetAllComments()
	if err != nil {
		log.Printf("ОШИБКА в GetAllComments: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve comments",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Comments": аllComments})
}
func (h *Handlers) CreateComment(ctx *gin.Context) {
	var comment models.Comment
	if err := ctx.BindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := h.DbPool.CreateComment(comment.UserID, comment.PostID, comment.Content); err != nil {
		log.Printf("ОШИБКА в CreateComment: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create comment",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
}
func (h *Handlers) GetAllDirects(ctx *gin.Context) {
	аllDirects, err := h.DbPool.GetAllDirects()
	if err != nil {
		log.Printf("ОШИБКА в GetAllDirects: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve directs",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Directs": аllDirects})
}
func (h *Handlers) CreateDirect(ctx *gin.Context) {
	var direct models.Direct
	if err := ctx.BindJSON(&direct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := h.DbPool.CreateDirect(direct.UserID, direct.FriendID); err != nil {
		log.Printf("ОШИБКА в CreateDirect: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create direct",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Direct created successfully"})
}
func (h *Handlers) GetAllMessages(ctx *gin.Context) {
	аllMessages, err := h.DbPool.GetAllMessages()
	if err != nil {
		log.Printf("ОШИБКА в GetAllMessages: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve messages",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Messages": аllMessages})
}
func (h *Handlers) CreateMessage(ctx *gin.Context) {
	var message models.Message
	if err := ctx.BindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := h.DbPool.CreateMessage(message.SenderID, message.ReceiverID, message.Content); err != nil {
		log.Printf("ОШИБКА в CreateMessage: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create message",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Message created successfully"})
}


func (h *Handlers) GetAllLikes(ctx *gin.Context) {
	аllLikes, err := h.DbPool.GetAllLikes()
	if err != nil {
		log.Printf("ОШИБКА в GetAllLikes: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve likes",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Likes": аllLikes})
}
func (h *Handlers) CreateLike(ctx *gin.Context) {
	var like models.Like
	if err := ctx.BindJSON(&like); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := h.DbPool.CreateLike(like.PostID, like.UserID); err != nil {
		log.Printf("ОШИБКА в CreateLike: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create like",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Like created successfully"})
}

func (h *Handlers) DeletePost(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
        return
    }
	user, err := h.DbPool.GetUserById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetUserById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve user",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"User": user})
}
func (h *Handlers) DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	user, err := h.DbPool.GetUserById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetUserById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve user",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"User": user})
}
func (h *Handlers) DeleteMessage(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	message, err := h.DbPool.GetMessageById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetMessageById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve message",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": message})
}
func (h *Handlers) DeleteComment(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	comment, err := h.DbPool.GetCommentById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetCommentById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve comment",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Comment": comment})
}
func (h *Handlers) DeleteLike(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	like, err := h.DbPool.GetLikeById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetLikeById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve like",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Like": like})
}
func (h *Handlers) DeleteDirect(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	direct, err := h.DbPool.GetDirectById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetDirectById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve direct",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Direct": direct})
}
func (h *Handlers) DeleteFriend(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	friend, err := h.DbPool.GetFriendById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetFriendById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve friend",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Friend": friend})
}
func (h *Handlers) UpdatePost(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	post, err := h.DbPool.GetPostById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetPostById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve post",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Post": post})
}
func (h *Handlers) GetPostById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	post, err := h.DbPool.GetPostById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetPostById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve post",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Post": post})
}
func (h *Handlers) GetLikeById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	like, err := h.DbPool.GetLikeById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetLikeById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve like",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Like": like})
}
func (h *Handlers) GetCountComments(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	count, err := h.DbPool.GetCountComments(id)
	if err != nil {
		log.Printf("ОШИБКА в GetCountComments: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve count",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Count": count})
}
func (h *Handlers) GetCountLikes(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	count, err := h.DbPool.GetCountLikes(id)
	if err != nil {
		log.Printf("ОШИБКА в GetCountLikes: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve count",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Count": count})
}
func (h *Handlers) GetCommentById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	comment, err := h.DbPool.GetCommentById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetCommentById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve comment",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Comment": comment})
}
func (h *Handlers) GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}
	user, err := h.DbPool.GetUserById(id)
	if err != nil {
		log.Printf("ОШИБКА в GetUserById: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve user",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"User": user})
}
func (h *Handlers) GetAllNotifications(ctx *gin.Context) {
	notifications, err := h.DbPool.GetAllNotifications()
	if err != nil {
		log.Printf("ОШИБКА в GetAllNotifications: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve notifications",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Notifications": notifications})
}
func (h *Handlers) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := h.DbPool.GetUserByUsername(username)
	if err != nil {
		log.Printf("ОШИБКА в GetUserByUsername: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve user",
			"details": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"User": user})
}
func (h *Handlers) GetPostsById()


	









