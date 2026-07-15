package database

import (
	"context"
	"semen_project/models"
)

func (s *Store) SaveRefreshToken(r models.RefreshToken) error {
	query := `
	INSERT INTO refresh_tokens (user_id, token_hash, expires_at, created_at)
	VALUES ($1, $2, $3, $4)
	`
	_, err := s.db.Exec(context.Background(), query, r.UserID, r.TokenHash, r.ExpiresAt, r.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) DeleteRefreshToken(token_id int) error {
	query := `DELETE FROM refresh_tokens WHERE id = $1`
	_, err := s.db.Exec(context.Background(), query, token_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) DeleteAllUserRefreshTokens(user_id int) error {
	query := `DELETE FROM refresh_tokens WHERE user_id = $1`
	_, err := s.db.Exec(context.Background(), query, user_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetRefreshToken(token_id int) (models.RefreshToken, error) {
	query := `
	SELECT id, user_id, token_hash, expires_at, created_at
	FROM refresh_tokens
	WHERE id = $1
	`
	var r models.RefreshToken
	err := s.db.QueryRow(context.Background(), query, token_id).Scan(
		&r.ID,
		&r.UserID,
		&r.TokenHash,
		&r.ExpiresAt,
		&r.CreatedAt,
	)
	if err != nil {
		return models.RefreshToken{}, err
	}
	return r, nil
}
func (s *Store) UpdateRefreshToken(r models.RefreshToken) error {
	query := `
	UPDATE refresh_tokens
	SET user_id = $1, token_hash = $2, expires_at = $3, created_at = $4
	WHERE id = $5
	`
	_, err := s.db.Exec(context.Background(), query, r.UserID, r.TokenHash, r.ExpiresAt, r.CreatedAt, r.ID)
	if err != nil {
		return err
	}
	return nil
}
