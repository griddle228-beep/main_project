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

func (u *UserStore)GetAllUsers() ([]models.User, error) {
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
		err := rows.Scan(&user.UserName, &user.FirstName, &user.LastName, &user.Password)
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