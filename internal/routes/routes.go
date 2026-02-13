package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)
func SetupRoutes(r *gin.Engine, h *controller.Handlers) {
	// GET запросы
	r.GET("/hello", controller.Hello)
	r.GET("/answer", controller.Answer)
	// POST запросы
	r.POST("/createuser", h.Create)
}