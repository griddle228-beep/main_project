package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupFollowRoutes(r *gin.RouterGroup, h *controller.Handlers) {
	r.("/followuser", h.)
	r.("/unfollowuser", h.)
	r.("/getallfollowers", h.)
	r.("/getallfollowing", h.)
	r.("/getcountfollowers", h.)
	r.("/getcountfollowing", h.)
}