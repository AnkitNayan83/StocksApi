// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, firstName, lastName, email, hashedPassword, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, firstname, lastname, email, hashedpassword, created_at, updated_at
`

type CreateUserParams struct {
	ID             uuid.UUID
	Firstname      string
	Lastname       sql.NullString
	Email          string
	Hashedpassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Hashedpassword,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Hashedpassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}