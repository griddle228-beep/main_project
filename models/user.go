package models

type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}
type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"post_id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}
type Post struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}
type Like struct {
	ID     int `json:"id"`
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
}
type Friend struct {
	ID       int `json:"id"`
	UserFirst   int `json:"user_first"`
	FriendSecond int `json:"user_second"`
}
type Chat struct {
	ID         int `json:"id"`
	UserFirst  int `json:"user_first"`
	UserSecond int `json:"user_second"`
}
type Message struct {
	ID         int    `json:"id"`
	ChatID     int    `json:"chat_id"`
	SenderID   int    `json:"sender_id"`
	Content    string `json:"content"`
	ReadMark   bool   `json:"mark_read"`
}
type following struct {
	ID     		int `json:"id"`
	FollowerId 	int `json:"follower_id"`
	UserID 		int `json:"user_id"`
}
type Notification struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}
type UserPublic struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Followers struct {
	ID     		int `json:"id"`
	FollowerID 	int `json:"follower_id"`
	UserID 		int `json:"user_id"`
}
