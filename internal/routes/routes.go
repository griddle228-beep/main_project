package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *controller.Handlers) {

	r.GET("/getallusers", h.GetAllUsers)


	// /auth
	r.POST("/authentication", h.Authentication)
	// /create_user
	r.POST("/createuser", h.Create)
	// /create_post
	r.POST("/createpost", h.CreatePost)
	r.POST("/deletepost", h.DeletePost)
	r.POST("/updatepost", h.UpdatePost)
	// /profile

	// /feed
	r.GET("/getallposts", h.GetAllPosts)
	r.GET("/getalllikes", h.GetAllLikes)
	r.GET("/getlikebyid/:id", h.GetLikeById)
	r.GET("/getallcomments", h.GetAllComments)
	r.GET("/getcommentbyid/:id", h.GetCommentById)
	r.GET("/getcountlikes/:id", h.GetCountLikes)
	r.POST("/createlike", h.CreateLike) 
	r.POST("/deletelike", h.DeleteLike) 
	r.POST("/deletecomment", h.DeleteComment) 
	r.POST("/createcomment", h.CreateComment) 
	// /profile
	r.GET("/getuserbyid/:id", h.GetUserById)
	// /activity

	// /explore
	r.GET("/getpostbyid/:id", h.GetPostById)
	r.GET("/userbeusername", h.GetUserByUsername)
	r.GET("/getpostsbyid/:id", h.GetPostsById)

	// /settings

	// / direct_messege

	// /friends
	r.GET("/friends/:id", h.GetAllFriends)	// /Notifications
	r.GET("/notifications", h.GetAllNotifications)
}
