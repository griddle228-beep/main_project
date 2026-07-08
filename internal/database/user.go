package database

import (
	"context"
	"semen_project/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStore struct {
	db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) *UserStore {
	return &UserStore{db: db}
}

func (u *UserStore) GetAllUsers() ([]models.User, error) {
	var users []models.User

	query := `
	SELECT * FROM users;
	`
	rows, err := u.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (s *UserStore) CreateUser(u models.User) (*models.User, error) {
	query := `
	INSERT INTO users (user_name, first_name, last_name, password)
	VALUES ($1, $2, $3, $4)
	RETURNING id, user_name, first_name, last_name, password;
	`
	createdUser := &models.User{}
	err := s.db.QueryRow(
		context.Background(),
		query,
		u.UserName, u.FirstName, u.LastName, u.Password,
	).Scan(
		&createdUser.ID,
		&createdUser.UserName,
		&createdUser.FirstName,
		&createdUser.LastName,
		&createdUser.Password,
	)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
func (s *UserStore) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, user_name, first_name, last_name, password FROM users WHERE user_name = $1`

	err := s.db.QueryRow(context.Background(), query, username).Scan(
		&user.ID,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (s *UserStore) GetAllFriends() ([]models.User, error) {
	var users []models.User

	query := `
    SELECT u.id, u.user_name, u.first_name, u.last_name 
    FROM users u
    JOIN friends f ON u.id = f.friend_id
    WHERE f.user_id = $1;
    `
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *UserStore) GetUserById(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, user_name, first_name, last_name, password FROM users WHERE id = $1`

	err := s.db.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (s *UserStore) AddFriend(user_id int, friend_id int) error {
	query := `INSERT INTO friends (user_id, friend_id) VALUES ($1, $2);`
	_, err := s.db.Exec(context.Background(), query, user_id, friend_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) DeleteFriend(user_id int, friend_id int) error {
	query := `DELETE FROM friends WHERE user_id = $1 AND friend_id = $2;`
	_, err := s.db.Exec(context.Background(), query, user_id, friend_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) GetFriendId(user_id int) ([]int, error) {
	var friend_id []int
	query := `SELECT friend_id FROM friends WHERE user_id = $1;`
	rows, err := s.db.Query(context.Background(), query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		friend_id = append(friend_id, id)
	}
	return friend_id, nil
}
func (s *UserStore) AddPost(user_id int, content string) error {
	query := `INSERT INTO posts (user_id, content) VALUES ($1, $2);`
	_, err := s.db.Exec(context.Background(), query, user_id, content)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	query := `SELECT id, user_id, content FROM posts;`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func (s *UserStore) GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment
	query := `SELECT id, post_id, user_id, content FROM comments;`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
func (s *UserStore) CreateComment(post_id int, user_id int, content string) error {
	query := `INSERT INTO comments (post_id, user_id, content) VALUES ($1, $2, $3);`
	_, err := s.db.Exec(context.Background(), query, post_id, user_id, content)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) DeleteComment(comment_id int) error {
	query := `DELETE FROM comments WHERE id = $1;`
	_, err := s.db.Exec(context.Background(), query, comment_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) GetUserPostsId(user_id int) ([]int, error) {
	var post_id []int
	query := `SELECT id FROM posts WHERE user_id = $1;`
	rows, err := s.db.Query(context.Background(), query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		post_id = append(post_id, id)
	}
	return post_id, nil
}
func (s *UserStore) CreateLike(post_id int, user_id int) error {
	query := `INSERT INTO likes (post_id, user_id) VALUES ($1, $2); ON CONFLICT DO NOTHING;`
	_, err := s.db.Exec(context.Background(), query, post_id, user_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) DeleteLike(post_id int, user_id int) error {
	query := `DELETE FROM likes WHERE post_id = $1 AND user_id = $2;`
	_, err := s.db.Exec(context.Background(), query, post_id, user_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) GetAllLikes() ([]models.Like, error) {
	var likes []models.Like
	query := `SELECT post_id, user_id FROM likes;`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var like models.Like
		err := rows.Scan(&like.PostID, &like.UserID)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	return likes, nil
}
func (s *UserStore) GetCountLikes(post_id int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM likes WHERE post_id = $1;`
	err := s.db.QueryRow(context.Background(), query, post_id).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (s *UserStore) GetCountComments(post_id int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM comments WHERE post_id = $1;`
	err := s.db.QueryRow(context.Background(), query, post_id).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (s *UserStore) GetCountPosts(user_id int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM posts WHERE user_id = $1;`
	err := s.db.QueryRow(context.Background(), query, user_id).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (s *UserStore) CreatePost(user_id int, content string) error {
	query := `INSERT INTO posts (user_id, content) VALUES ($1, $2);`
	_, err := s.db.Exec(context.Background(), query, user_id, content)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) DeletePost(post_id int) error {
	query := `DELETE FROM posts WHERE id = $1;`
	_, err := s.db.Exec(context.Background(), query, post_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) GetAllDirects() ([]models.Direct, error) {
	var directs []models.Direct
	query := `SELECT id, user_id, friend_id FROM directs;`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var direct models.Direct
		err := rows.Scan(&direct.ID, &direct.UserID, &direct.FriendID)
		if err != nil {
			return nil, err
		}
		directs = append(directs, direct)
	}
	return directs, nil
}
func (s *UserStore) CreateDirect(user_id int, friend_id int) error {
	query := `INSERT INTO directs (user_id, friend_id) VALUES ($1, $2);`
	_, err := s.db.Exec(context.Background(), query, user_id, friend_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) DeleteDirect(direct_id int) error {
	query := `DELETE FROM directs WHERE id = $1;`
	_, err := s.db.Exec(context.Background(), query, direct_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) GetUserByDirectID(direct_id int) (int, error) {
	var user_id int
	query := `SELECT user_id FROM directs WHERE id = $1;`
	err := s.db.QueryRow(context.Background(), query, direct_id).Scan(&user_id)
	if err != nil {
		return 0, err
	}
	return user_id, nil
}
func (s *UserStore) GetFriendByDirectID(direct_id int) (int, error) {
	var friend_id int
	query := `SELECT friend_id FROM directs WHERE id = $1;`
	err := s.db.QueryRow(context.Background(), query, direct_id).Scan(&friend_id)
	if err != nil {
		return 0, err
	}
	return friend_id, nil
}
func (s *UserStore) GetDirectID(user_id int, friend_id int) (int, error) {
	var direct_id int
	query := `SELECT id FROM directs WHERE user_id = $1 AND friend_id = $2;`
	err := s.db.QueryRow(context.Background(), query, user_id, friend_id).Scan(&direct_id)
	if err != nil {
		return 0, err
	}
	return direct_id, nil
}
func (s *UserStore) ExploreUsers(user_id int) ([]models.User, error) {
	var users []models.User
	query := `SELECT id, username FROM users WHERE id != $1;`
	rows, err := s.db.Query(context.Background(), query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.UserName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (s *UserStore) ExploreGlobalPosts() ([]models.Post, error) {
	var posts []models.Post
	query := `SELECT id, user_id, content FROM posts;`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func (s *UserStore) GetAllSubscriptions() ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	query := `SELECT id, user_id FROM subscriptions;`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var subscription models.Subscription
		err := rows.Scan(&subscription.ID, &subscription.UserID)
		if err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}
func (s *UserStore) CreateSubscription(user_id int) error {
	query := `INSERT INTO subscriptions (user_id) VALUES ($1);`
	_, err := s.db.Exec(context.Background(), query, user_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) DeleteSubscription(subscription_id int) error {
	query := `DELETE FROM subscriptions WHERE id = $1;`
	_, err := s.db.Exec(context.Background(), query, subscription_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) GetUserBySubscriptionID(subscription_id int) (int, error) {
	var user_id int
	query := `SELECT user_id FROM subscriptions WHERE id = $1;`
	err := s.db.QueryRow(context.Background(), query, subscription_id).Scan(&user_id)
	if err != nil {
		return 0, err
	}
	return user_id, nil
}
func (s *UserStore) GetAllSubscribers() ([]models.Subscriber, error) {
	var subscribers []models.Subscriber
	query := `SELECT id, user_id FROM subscribers;`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var subscriber models.Subscriber
		err := rows.Scan(&subscriber.ID, &subscriber.UserID)
		if err != nil {
			return nil, err
		}
		subscribers = append(subscribers, subscriber)
	}
	return subscribers, nil

}
func (s *UserStore) CreateSubscriber(user_id int) error {
	query := `INSERT INTO subscribers (user_id) VALUES ($1);`
	_, err := s.db.Exec(context.Background(), query, user_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) DeleteSubscriber(subscriber_id int) error {
	query := `DELETE FROM subscribers WHERE id = $1;`
	_, err := s.db.Exec(context.Background(), query, subscriber_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) GetUserBySubscriberID(subscriber_id int) (int, error) {
	var user_id int
	query := `SELECT user_id FROM subscribers WHERE id = $1;`
	err := s.db.QueryRow(context.Background(), query, subscriber_id).Scan(&user_id)
	if err != nil {
		return 0, err
	}
	return user_id, nil
}
func (s *UserStore) GetActivity(user_id int) ([]models.Post, error) {
	var posts []models.Post
	query := `SELECT id, user_id, content FROM posts WHERE user_id = $1 OR id IN (SELECT post_id FROM comments WHERE user_id = $1);`
	rows, err := s.db.Query(context.Background(), query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func (s *UserStore) GetAllPostsByUserID(user_id int) ([]models.Post, error) {
	var posts []models.Post
	query := `SELECT id, user_id, content FROM posts WHERE user_id = $1;`
	rows, err := s.db.Query(context.Background(), query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func (s *UserStore) GetAllMessages() ([]models.Message, error) {
	var messages []models.Message
	query := `SELECT id, user_id, content FROM messages;`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.ID, &message.SenderID, &message.Content)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
func (s *UserStore) CreateMessage(sender_id int, receiver_id int, content string) error {
	query := `INSERT INTO messages (sender_id, receiver_id, content) VALUES ($1, $2, $3);`
	_, err := s.db.Exec(context.Background(), query, sender_id, receiver_id, content)
	if err != nil {
		return err
	}
	return nil
}
func (s *UserStore) DeleteMessage(message_id int) error {
	query := `DELETE FROM messages WHERE id = $1;`
	_, err := s.db.Exec(context.Background(), query, message_id)
	if err != nil {
		return err
	}
	return nil
}
func (s * UserStore) GetMessageById(message_id int) (models.Message, error) {
	var message models.Message
	query := `SELECT id, sender_id, receiver_id, content FROM messages WHERE id = $1;`
	err := s.db.QueryRow(context.Background(), query, message_id).Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Content)
	if err != nil {
		return models.Message{}, err
	}
	return message, nil
}
func (s * UserStore) GetCommentById(comment_id int) (models.Comment, error) {
	var comment models.Comment
	query := `SELECT id, user_id, post_id, content FROM comments WHERE id = $1;`
	err := s.db.QueryRow(context.Background(), query, comment_id).Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content)
	if err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}
func (s * UserStore) GetAllCommentsByPostID(post_id int) ([]models.Comment, error) {
	var comments []models.Comment
	query := `SELECT id, user_id, post_id, content FROM comments WHERE post_id = $1;`
	rows, err := s.db.Query(context.Background(), query, post_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
func (s * UserStore) GetLikeById(like_id int) (models.Like, error) {
	var like models.Like
	query := `SELECT id, user_id, post_id FROM likes WHERE id = $1;`
	err := s.db.QueryRow(context.Background(), query, like_id).Scan(&like.ID, &like.UserID, &like.PostID)
	if err != nil {
		return models.Like{}, err
	}
	return like, nil
}
func (s * UserStore) GetDirectById(direct_id int) (models.Direct, error) {
	var direct models.Direct
	query := `SELECT id, user_id, receiver_id FROM direct WHERE id = $1;`
	err := s.db.QueryRow(context.Background(), query, direct_id).Scan(&direct.ID, &direct.UserID, &direct.FriendID)
	if err != nil {
		return models.Direct{}, err
	}
	return direct, nil
}
func (s * UserStore) GetFriendById(friend_id int) (models.Friend, error) {
	var friend models.Friend
	query := `SELECT id, user_id, friend_id FROM friends WHERE id = $1;`
	err := s.db.QueryRow(context.Background(), query, friend_id).Scan(&friend.ID, &friend.UserID, &friend.FriendID)
	if err != nil {
		return models.Friend{}, err
	}
	return friend, nil
}
func (s * UserStore) GetPostById(post_id int) (models.Post, error) {
	var post models.Post
	query := `SELECT id, user_id, content FROM posts WHERE id = $1;`
	err := s.db.QueryRow(context.Background(), query, post_id).Scan(&post.ID, &post.UserID, &post.Content)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}
func (s * UserStore) GetAllNotifications() ([]models.Notification, error) {
	var notifications []models.Notification
	query := `SELECT id, user_id, content FROM notifications;`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var notification models.Notification
		err := rows.Scan(&notification.ID, &notification.UserID, &notification.Content)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	return notifications, nil
}
func (s * UserStore) GetAllPostsById(user_id int) ([]models.Post, error) {
	var posts []models.Post
	query := `SELECT id, user_id, content FROM posts WHERE user_id = $1;`
	rows, err := s.db.Query(context.Background(), query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
	