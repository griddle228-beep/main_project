package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *controller.Handlers) {
	// GET запросы

	r.GET("/getallusers", h.GetAllUsers)
	// POST запросы

	// /auth
	r.POST("/authentication", h.Authentication)
	// /create_user
	r.POST("/createuser", h.Create)
	// /create_post
	r.POST("/createpost", h.CreatePost)
	// /profile
	
	// /feed

	// /activity

	// /explore
	
	// /settings

	// / direct_messege

	// /friends
	r.GET("/friends", h.GetAllFriends)
}
