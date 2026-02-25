

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    password VARCHAR(255) NOT NULL
);
INSERT INTO users (user_name, first_name, last_name, password) VALUES ('admin', 'Admin', 'Adminov', '1111');

