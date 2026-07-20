package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup, h *controller.Handlers) {
// user
	r.("/getallusers", h.)
	r.("/searchusers", h.)
	r.("/updatepassword", h.)
	r.("/updateuser", h.)
}
