package database

import (
	"context"
	"semen_project/models"

)

func (s *Store) CreateComment(c models.Comment) error{
	query := `
	INSERT INTO comments (post_id, user_id, content)
	VALUES ($1, $2, $3)
	`
	_, err := s.db.Exec(context.Background(),query, c.PostID, c.UserID, c.Content)
	if err != nil {
		return  err
	}
	return nil
}
func (s *Store) DeleteComment(comment_id int) error {	
	query := `
	DELETE FROM comments WHERE user_id = $1 AND post_id = $2
	`
	_, err := s.db.Exec(context.Background(),query, comment_id)
	if err != nil {
		return  err
	}
	return nil
}
func (s *Store) UpdateComment(comment_id int,content string) error {
	query := `
	UPDATE comments
	SET content = $1
	WHERE id = $1
	`
	_, err := s.db.Exec(context.Background(),query,comment_id, content)
	if err != nil {
		return  err
	}
	return nil
}
func (s *Store) GetAllPostComments(post_id int) ([]models.Comment, error) {
	var comments []models.Comment
	query := `
	SELECT post_id, user_id, content
	FROM comments
	WHERE post_id = $1
	`
	rows, err := s.db.Query(context.Background(),query, post_id)
	if err != nil {
		return  nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.PostID, &comment.UserID, &comment.Content)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
    return nil, err
	}	
	return comments, nil
}