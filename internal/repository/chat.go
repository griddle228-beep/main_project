package repository

import (
	"context"
	"semen_project/internal/models"
)


func (s *Store) CreateChat(user_first_id int, user_seconf_id int) error {
	query := `
	INSERT INTO chats (user_first, user_second)
	VALUES ($1, $2)
	`
	_, err := s.db.Exec(context.Background(), query, user_first_id, user_seconf_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetAllUserChats(user_id int) ([]models.Chat, error) {
	var chats []models.Chat
	query := `
	SELECT id, user_first, user_second
	FROM chats
	WHERE user_first = $1 OR user_second = $1
	`
	rows, err := s.db.Query(context.Background(), query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var chat models.Chat
		err := rows.Scan(&chat.ID, &chat.UserFirst, &chat.UserSecond)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	return chats, nil
}
func (s *Store) DeleteChat(chat_id int) error {
	query := `
	DELETE FROM chats
	WHERE id = $1
	`
	_, err := s.db.Exec(context.Background(), query, chat_id)
	if err != nil {
		return err
	}
	return nil
}