package routes

import (
	controller "semen_project/internal/controllers"
	"semen_project/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *controller.Handlers, secret string) {

    api := r.Group("/api")

    public := api.Group("")
    private := api.Group("")

    private.Use(middleware.AuthMiddleware(secret))

    SetupPublicAuthenticationRoutes(public, h)

    SetupPrivateAuthenticationRoutes(private, h)

    SetupUserRoutes(private, h)
    SetupPostRoutes(private, h)
    SetupMessageRoutes(private, h)
    SetupLikeRoutes(private, h)
    SetupFriendRoutes(private, h)
    SetupFollowRoutes(private, h)
    SetupCommentRoutes(private, h)
    SetupChatRoutes(private, h)
}