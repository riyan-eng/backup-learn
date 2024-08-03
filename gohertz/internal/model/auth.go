package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id        string         `db:"id"`
	Email     sql.NullString `db:"email"`
	Username  sql.NullString `db:"username"`
	Password  sql.NullString `db:"password"`
	IsActive  sql.NullBool   `db:"is_active"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

type UserData struct {
	Id       string         `db:"id"`
	UserId   sql.NullString `db:"user_id"`
	RoleCode sql.NullString `db:"role_code"`
}
