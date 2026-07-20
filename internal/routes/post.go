package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(r *gin.RouterGroup, h *controller.Handlers) {
	r.("/createpost", h.)
	r.("/getfeed", h.)
	r.("/getpostbyid", h.)
	r.("/getalluserposts", h.)
	r.("/getallposts", h.)
	r.("/deletepost", h.)
	r.("/updatepost", h.)
}