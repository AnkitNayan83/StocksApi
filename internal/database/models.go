// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Stock struct {
	ID            uuid.UUID
	Companyname   string
	Valueperstock float64
	Quantity      int32
	Ownerid       uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Stockbuy struct {
	ID        uuid.UUID
	Userid    uuid.UUID
	Stockid   uuid.UUID
	Quantity  int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID             uuid.UUID
	Firstname      string
	Lastname       sql.NullString
	Email          string
	Hashedpassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
