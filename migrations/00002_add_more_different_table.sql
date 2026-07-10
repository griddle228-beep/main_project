-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS friend (
    id SERIAL PRIMARY KEY,
    user_first INT NOT NULL,
    user_second INT NOT NULL,
    FOREIGN KEY (user_first) REFERENCES users(id),
    FOREIGN KEY (user_second) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS post (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    content VARCHAR(10000) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS comment (
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL,
    user_id INT NOT NULL,
    content VARCHAR(10000) NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS like (
    id SERIAL PRIMARY KEY,
    post_id INT NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS chat (
    id SERIAL PRIMARY KEY,
    user_first INT NOT NULL, 
    user_second INT NOT NULL,
    FOREIGN KEY (user_first) REFERENCES users(id),
    FOREIGN KEY (user_second) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS message (
    id SERIAL PRIMARY KEY,
    chat_id INT NOT NULL,
    sender_id INT NOT NULL,
    content VARCHAR(10000) NOT NULL,
    mark_read BOOLEAN,
    FOREIGN KEY (chat_id) REFERENCES chat(id),
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (receiver_id) REFERENCES users(id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS chat, message, like, comment, post, friend;
-- +goose StatementEnd
