package database

import (
	"context"
	"semen_project/models"

)

func (s *Store) CreatePost(p models.Post) (*models.Post, error)  {
	post := &models.Post{}
	query := ` INSERT INTO posts (user_id, content)
	VALUES ($1, $2)
	RETURNING id, user_id, content;
	`
	err := s.db.QueryRow(context.Background(), query, p.UserID, p.Content).Scan(
		&post.ID,
		&post.UserID,
		&post.Content,
	)
	if err != nil {
		return nil, err
	}
	return post, err 
}
func (s *Store) GetPostById(user_id int) (models.Post, error)  {
	var post models.Post
	query := `SELECT id, user_id, content FROM posts WHERE id = $1 `
	err := s.db.QueryRow(context.Background(), query, user_id).Scan(&post.ID, &post.UserID, &post.Content)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}
func (s *Store) DeletePost(post_id int) error {
	query := `DELETE FROM posts WHERE id = $1 `
	_, err := s.db.Exec(context.Background(), query, post_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetAllPosts() ([]models.Post, error)  {
	var posts []models.Post
	query := ` SELECT Id, user_id, content From posts`
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
	if err := rows.Err(); err != nil {
    return nil, err
	}	
	return posts, nil
}
func (s *Store) GetAllUserPosts(user_id int) ([]models.Post, error) {
	var posts []models.Post
	query := `SELECT id, user_id, content FROM posts WHERE user_id = $1`
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
	if err := rows.Err(); err != nil {
    return nil, err
	}	
	return posts, nil
}
func (s *Store) UpdatePost(p models.Post) (*models.Post, error)  {
	updatedpost := &models.Post{}
	query := `
	UPDATE posts
	SET user_id = $2, content = $3
	WHERE id = $1
	RETURNING id, user_id, content
	`
	err := s.db.QueryRow(context.Background(), query, p.ID, p.UserID, p.Content).Scan(
		&updatedpost.ID,
		&updatedpost.UserID,
		&updatedpost.Content,
	)
	if err != nil {
		return nil, err
	}
	return updatedpost, nil
}
// for feed
func (s *Store) GetAllFriendsPosts(user_id int) ([]models.Post, error) {
	var posts []models.Post
	query := `
	SELECT p.id, p.user_id, p.content
	FROM posts p
	JOIN friends f
	ON (
 	   (f.user_first = $1 AND p.user_id = f.user_second)
  	  OR
  	  (f.user_second = $1 AND p.user_id = f.user_first)
	);	
	`
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
	if err := rows.Err(); err != nil {
    return nil, err
	}	
	return posts, nil
}
func (s *Store) GetAllNotFriendsPosts(id int) ([]models.Post, error) {
	var posts []models.Post
	query := `
	SELECT p.id, p.user_id, p.content
	FROM posts p
	WHERE p.user_id <> $1
	AND NOT EXISTS (
	    SELECT 1
	    FROM friends f
	    WHERE
	        (f.user_first = $1 AND f.user_second = p.user_id)
	        OR
	        (f.user_second = $1 AND f.user_first = p.user_id)
);
	`
	rows, err := s.db.Query(context.Background(), query, id)
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
	if err := rows.Err(); err != nil {
    return nil, err
	}	
	return posts, nil
}