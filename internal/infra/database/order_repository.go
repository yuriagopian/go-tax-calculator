package database

import "database/sql"

type OrderRepository struct {
	Db *sql.DB
}
