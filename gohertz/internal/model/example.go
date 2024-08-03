package model

import "database/sql"

type Example struct {
	Id        string
	Name      sql.NullString  `db:"id"`
	Detail    sql.NullString  `db:"name"`
	Count     sql.NullFloat64 `db:"count"`
	CreatedAt sql.NullTime    `db:"created_at"`
	UpdatedAt sql.NullTime    `db:"udpated_at"`
}
