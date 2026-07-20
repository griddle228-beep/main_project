package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupChatRoutes(r *gin.RouterGroup, h *controller.Handlers) {
	r.("/createchat", h.)
	r.("/getalluserchats", h.)
	r.("/deletechat", h.)
}