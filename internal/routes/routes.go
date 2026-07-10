package routes

import (
	controller "semen_project/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *controller.Handlers) {
// user
	r.("/getallusers", h.)
	r.("/searchusers", h.)
	r.("/follow", h.)
	r.("/getuserbyid", h.)
	r.("/getuserbynickname", h.)
	r.("/getuserbyfirstname", h.)
	r.("/getuserbylastname", h.)
	r.("/getallfollowed", h.)
	r.("/getallfollowings", h.)
	r.("/getallfriends", h.)
// post
	r.("/createpost", h.)
	r.("/getallposts", h.)
	r.("/getpostbyid", h.)
	r.("/getalluserposts", h.)
	r.("/deletepost", h.)
	r.("/updatepost", h.)
// like
	r.("/likepost", h.)
	r.("/deletelike", h.)
	r.("/countlikes", h.)
	r.("/getalllikes", h.)
// comment
	r.("/createcomment", h.)
	r.("/deletecomment", h.)
	r.("/updatecomment", h.)
	r.("/getallcomments", h.)
	r.("/getcountcomments", h.)
// chat
	r.("/createchat", h.)
	r.("/sendmessage", h.)
	r.("/getallchats", h.)
	r.("/deletemessage", h.)
	r.("/deletechat", h.)
	r.("/getallmessages", h.)
	r.("/getmarkread", h.)
	r.("/getcountnotreadmessage", h.)
// notification
	r.("/getallnotifications", h.)
	r.("/getnotification", h.)
	r.("/createnotification", h.)
	r.("/deletenotification", h.)
// authentication
	r.("/register", h.)
	r.("/login", h.)
	r.("/refreshtoken", h.)
	r.("/logout", h.)






	// /create_user

	// /create_post
	r.POST("/createpost/:id", h.CreatePost)
	// /profile

	// /feed

	// /profile
	// /activity

	// /explore


	// /settings

	// / Chats

	// /friends

	// /Notifications
}
