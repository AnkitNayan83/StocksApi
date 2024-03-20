-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    firstName TEXT NOT NULL,
    lastName TEXT,
    email VARCHAR(30) NOT NULL,
    hashedPassword VARCHAR(15) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down

DROP TABLE users;