package repository

import (
	"context"
	"semen_project/internal/models"

)

func (s *Store) LikePost(post_id int, user_id int) error {
	query := `
	INSERT INTO likes (post_id, user_id)
	VALUES ($1, $2)
	`
	_, err := s.db.Exec(context.Background(), query, post_id, user_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) DeleteLike(like_id int) error {
	query := `
	DELETE FROM likes
	WHERE id = $1
	`
	_, err := s.db.Exec(context.Background(), query, like_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetAllLUserLikes(user_id int) ([]models.Like, error) {
	var likes []models.Like
	query := `
	SELECT id, post_id, user_id
	FROM likes
	WHERE id = $1
	`
	rows, err := s.db.Query(context.Background(),query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var like models.Like
		err := rows.Scan(&like.ID, &like.PostID, &like.UserID)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	if err := rows.Err(); err != nil {
    return nil, err
	}	
	return likes, nil
}
func (s *Store) GetAllLPostLikes(post_id int) ([]models.Like, error)  {
	var likes []models.Like
	query := `
	SELECT id, post_id, user_id
	FROM likes
	WHERE post_id = $1
	`
	rows, err := s.db.Query(context.Background(),query, post_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var like models.Like
		err := rows.Scan(&like.ID, &like.PostID, &like.UserID)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}
	if err := rows.Err(); err != nil {
    return nil, err
	}	
	return likes, nil
}