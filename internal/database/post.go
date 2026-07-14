package database

import (
	"context"
	"semen_project/models"

)

func (s *Store) CreatePost(u models.Post) (*models.Post, error)  {
	post := &models.Post{}
	query := ` INSERT INTO post (user_id, content)
	VALUES ($1, $2)
	RETURNING id, user_id, content;
	`
	err := s.db.QueryRow(context.Background(), query, u.UserID, u.Content).Scan(
		&post.ID,
		&post.UserID,
		&post.Content,
	)
	if err != nil {
		return nil, err
	}
	return post, err 
}
func (s *Store) GetPostById(id int) (models.Post, error)  {
	var post models.Post
	query := `SELECT id, user_id, content FROM post WHERE id = $1 `
	err := s.db.QueryRow(context.Background(), query, id).Scan(&post.ID, &post.UserID, &post.Content)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}
func (s *Store) DeletePost(id int) error {
	query := `SELECT id, user_id, content FROM post WHERE id = $1 `
	_, err := s.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetAllPosts() ([]models.Post, error)  {
	var posts []models.Post
	query := ` SELECT Id, user_id, content From post`
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
func (s *Store) GetAllUserPosts(user_id int) ([]models.Post, error) {
	var posts []models.Post
	query := `SELECT id, user_id, content FROM post WHERE user_id = $1`
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
func (s *Store) UpdatePost(p models.Post) (*models.Post, error)  {
	updatedpost := &models.Post{}
	query := `
	UPDATE post
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
func (s *Store) GetAllFriendPosts(id int) ([]models.Post, error) {
	var posts []models.Post
	query := `
	SELECT p.id, p.user_id, p.content
	FROM post p
	JOIN friend f
	ON (
 	   (f.user_first = $1 AND p.user_id = f.user_second)
  	  OR
  	  (f.user_second = $1 AND p.user_id = f.user_first)
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
	return posts, nil
}
func (s *Store) GetAllNotFriendPosts(id int) ([]models.Post, error) {
	var posts []models.Post
	query := `
	SELECT p.id, p.user_id, p.content
	FROM post p
	WHERE p.user_id <> $1
	AND NOT EXISTS (
	    SELECT 1
	    FROM friend f
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
	return posts, nil
}