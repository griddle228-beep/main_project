package repository

import (
	"context"
	"semen_project/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(u models.User) (*models.UserPublic, error) {
	query := `
	INSERT INTO users (user_name, first_name, last_name, password)
	VALUES ($1, $2, $3, $4)
	RETURNING id, user_name, first_name, last_name;
	`
	createdUser := &models.UserPublic{}
	err := s.db.QueryRow(
		context.Background(),
		query,
		u.UserName, u.FirstName, u.LastName, u.Password,
	).Scan(
		&createdUser.ID,
		&createdUser.UserName,
		&createdUser.FirstName,
		&createdUser.LastName,
	)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
func (s *Store) GetUserById(id int) (models.User, error) {
	var user models.User
	query := `SELECT id, user_name, first_name, last_name
	FROM users
	WHERE id = $1
	`
	err := s.db.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (s *Store) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	query := `SELECT id, user_name, first_name, last_name
	FROM users
	WHERE username = $1
	`
	err := s.db.QueryRow(context.Background(), query, username).Scan(
		&user.ID,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (s *Store) GetAllUsers() ([]models.UserPublic, error) {
	var users []models.UserPublic
	query := `SELECT id, nick_name, first_name, last_name FROM users`
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.UserPublic
		err := rows.Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
func (s *Store) UpdateUser(u models.User) (*models.UserPublic, error) {
	updateduser := &models.UserPublic{}
	query := `UPDATE users 
	SET user_name = $2, first_name = $3, last_name = $4
	WHERE id $1 
	RETURNING id, user_name, first_name, last_name;	
	`
	err := s.db.QueryRow(context.Background(), query, u.ID, u.UserName, u.FirstName, u.LastName).Scan(
		&updateduser.ID,
		&updateduser.UserName,
		&updateduser.FirstName,
		&updateduser.LastName,
	)
	if err != nil {
		return nil, err
	}
	return updateduser, nil
}
func (s *Store) ChangePassword(u models.User) error {
	query := `
	UPDATE users
	SET password = $2
	WHERE id = $1
	`
	_, err := s.db.Exec(context.Background(), query, u.ID, u.Password)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := s.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetUsersByFirstName(firstname string) ([]models.UserPublic, error) {
	var users []models.UserPublic
	query := `
	SELECT id, user_name, first_name, last_name
	FROM users
	WHERE first_name = $1
	`
	rows, err := s.db.Query(context.Background(), query, firstname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.UserPublic
		err := rows.Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
func (s *Store) GetUsersByLastName(lastname string) ([]models.UserPublic, error) {
	var users []models.UserPublic
	query := `
	SELECT id, user_name, first_name, last_name
	FROM users
	WHERE last_name = $1
	`
	rows, err := s.db.Query(context.Background(), query, lastname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.UserPublic
		err := rows.Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
func (s *Store) GetUsersByFullName(f string, l string) ([]models.UserPublic, error) {
	var users []models.UserPublic
	query := `
	SELECT id, user_name, first_name, last_name
	FROM users
	WHERE first_name = $1 AND last_name = $2
	`
	rows, err := s.db.Query(context.Background(), query, f, l)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.UserPublic
		err := rows.Scan(&user.ID, &user.UserName, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
