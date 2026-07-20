package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupMessageRoutes(r *gin.RouterGroup, h *controller.Handlers) {
	r.("/sendmessage", h.)
	r.("/updatemessage", h.)
	r.("/deletemessage", h.)
	r.("/getallmessages", h.)
	r.("/updatemarkreadtoread", h.)
	r.("/getcountnotreadmessage", h.)
}