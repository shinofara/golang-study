package infrastructure

import (
	"database/sql"
	"time"
)

const dsn = "root@tcp(db)/codetest"

// DB database interface
type DB struct {
	conn *sql.DB
}

// NewDB is DB constructor.
func NewDB() (*DB, error) {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db := DB{conn: conn}

	conn.SetConnMaxLifetime(10 * time.Second)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return &db, nil
}

// Open returns the database connection.
func (d *DB) Open() *sql.DB {
	return d.conn
}
