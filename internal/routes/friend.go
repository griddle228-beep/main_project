package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupFriendRoutes(r *gin.RouterGroup, h *controller.Handlers) {
	r.("/getcountfriends", h.)
	r.("/getallfriends", h.)
}