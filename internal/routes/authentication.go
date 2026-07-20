package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPublicAuthenticationRoutes(r *gin.RouterGroup, h *controller.Handlers) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/refresh", h.Refresh)
}
func SetupPrivateAuthenticationRoutes(r *gin.RouterGroup, h *controller.Handlers) {
	r.POST("/logout", h.Logout)
	r.POST("/logout/all-devices", h.LogoutAllDevicesExceptThis)
}