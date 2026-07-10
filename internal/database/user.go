package database

import (
	"context"
	"semen_project/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

// user
func (u *Store) CreateUser() {
}
func (u *Store) GetUserById() {
}
func (u *Store) GetUserByUsername() {
}
func (u *Store) UpdateUser() {
}
func (u *Store) ChangePassword() {
}
func (u *Store) DeleteUser() {
}
func (u *Store) SearchUsers() {
}
func (u *Store) FollowUser() {
}
func (u *Store) UnFollowUser() {
}
func (u *Store) DeleteFriend() {
}
func (u *Store) GetAllFriends() {
}
func (u *Store) GetAllFollowers() {
}
func (u *Store) GetAllFollowings() {
}

// like
func (u *Store) LikePost() {
}
func (u *Store) DeleteLike() {
}
func (u *Store) Countlikes() {
}
func (u *Store) AllLikes() {
}

// comment
func (u *Store) CreateComment() {
}
func (u *Store) DeleteComment() {
}
func (u *Store) UpdateComment() {
}
func (u *Store) GetAllComments() {
}
func (u *Store) GetCountComments() {
}

// chat
func (u *Store) CreateChat() {
}
func (u *Store) SendMessage() {
}
func (u *Store) GetAllChats() {
}
func (u *Store) DeleteMessage() {
}
func (u *Store) DeleteChat() {
}
func (u *Store) GetAllMessages() {
}
func (u *Store) GetMarkRead() {
}

// notification
func (u *Store) GetAllNotifications() {
}
func (u *Store) GetNotification() {
}
func (u *Store) CreateNotification() {
}
func (u *Store) DeleteNotification() {
}

// authentication
func (u *Store) SaveRefreshToken() {
}
func (u *Store) DeleteRefreshToken() {
}
func (u *Store) GetRefreshToken() {
}





















func (u *Store) GetAllUsers() ([]models.User, error) {
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
func (s *Store) CreateUser(u models.User) (*models.User, error) {
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
func (s *Store) GetUserByUsername(username string) (*models.User, error) {
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
func (s *Store) GetUserById(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, user_name, first_name, last_name, password FROM users WHERE id = $1`

	err := s.db.QueryRow(context.Background(), query, id).Scan(
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
func (s *Store) AddFriend(user_id int, friend_id int) error {
	query := `INSERT INTO friends (user_id, friend_id) VALUES ($1, $2);`
	_, err := s.db.Exec(context.Background(), query, user_id, friend_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) DeleteFriend(user_id int, friend_id int) error {
	query := `DELETE FROM friends WHERE user_id = $1 AND friend_id = $2;`
	_, err := s.db.Exec(context.Background(), query, user_id, friend_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetFriendId(user_id int) ([]int, error) {
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
func (s *Store) AddPost(user_id int, content string) error {
	query := `INSERT INTO posts (user_id, content) VALUES ($1, $2);`
	_, err := s.db.Exec(context.Background(), query, user_id, content)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetAllPosts() ([]models.Post, error) {
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
