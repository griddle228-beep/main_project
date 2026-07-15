package database

import (
	"context"
	"semen_project/models"
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
func (s *Store) SendMessage(m models.Message) error {
	query := `
	INSERT INTO messages (chat_id, sender_id, content)
	VALUES ($1, $2, $3)
	`
	_, err := s.db.Exec(context.Background(), query, m.ChatID, m.SenderID, m.Content)
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
func (s *Store) UpdateMessage(message_id int, content string) error {
	query := `
	UPDATE messages
	SET content = $2
	WHERE id = $1
	`
	_,err := s.db.Exec(context.Background(), query, message_id, content)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) DeleteMessage(message_id int) error {
	query := `
	DELETE FROM messages
	WHERE id = $1
	`
	_, err := s.db.Exec(context.Background(), query, message_id)
	if err != nil {
		return err
	}
	return nil
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
func (s *Store) GetAllChatMessages(chat_id int) ([]models.Message, error) {
	var messages []models.Message
	query := `
	SELECT id, chat_id, sender_id, content, mark_read
	FROM messages
	WHERE chat_id = $1
	`
	rows, err := s.db.Query(context.Background(), query, chat_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.ID, &message.ChatID, &message.SenderID, &message.Content, &message.MarkRead)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
func (s *Store) UnReadToRead(message_id int) error {
	query := `
	UPDATE message
	SET mark_read = true
	WHERE id = $1
	`
	_,err := s.db.Exec(context.Background(), query, message_id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) GetMessageById(message_id int) (models.Message, error) {
	var message models.Message
	query := `
	SELECT id, chat_id, sender_id, content, mark_read
	FROM messages
	WHERE id = $1
	`
	err := s.db.QueryRow(context.Background(), query, message_id).Scan(&message.ID,
		 &message.ChatID, &message.SenderID, &message.Content, &message.MarkRead)
		if err != nil {
			return models.Message{}, err
		}
	return message, nil
}
