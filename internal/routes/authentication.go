package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPublicAuthenticationRoutes(r *gin.RouterGroup, h *controller.Handlers) {
	r.("/register", h.)
	r.("/login", h.)
	r.("/refresh", h.)
}
func SetupPrivateAuthenticationRoutes(r *gin.RouterGroup, h *controller.Handlers) {
	r.("/logout", h.)
}