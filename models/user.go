package models
type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}
type Comment struct {
	ID        int    `json:"id"`
	PostID    int    `json:"post_id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
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
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
}
type Message struct {
	ID        int    `json:"id"`
	SenderID  int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Content   string `json:"content"`
}
type DirectMessage struct {
	ID        int    `json:"id"`
	SenderID  int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Content   string `json:"content"`
}
type Direct struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
}
type Subscription struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}
type Subscriber struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}
type Notification struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}