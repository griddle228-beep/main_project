package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCommentRoutes(r *gin.RouterGroup, h *controller.Handlers) {
	r.("/createcomment", h.)
	r.("/deletecomment", h.)
	r.("/updatecomment", h.)
	r.("/getallpostcomments", h.)
	r.("/getcountcomments", h.)
}