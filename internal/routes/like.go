package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupLikeRoutes(r *gin.RouterGroup, h *controller.Handlers) {
	r.("/likepost", h.)
	r.("/deletelike", h.)
	r.("/countlikes", h.)
	r.("/getallpostlikes", h.)
	r.("/getalluserlikes", h.)
}