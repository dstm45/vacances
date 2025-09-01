-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    username TEXT,
    email TEXT,
    password_hash TEXT
);

CREATE TABLE IF NOT EXISTS notes (
    id BIGSERIAL PRIMARY KEY,
    uuid uuid NOT NULL DEFAULT gen_random_uuid(),
    title TEXT UNIQUE,
    author_id INTEGER,
    FOREIGN KEY(author_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE notes;
DROP TABLE users;
