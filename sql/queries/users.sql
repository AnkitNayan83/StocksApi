-- name: CreateUser :one
INSERT INTO users (id, firstName, lastName, email, hashedPassword, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

