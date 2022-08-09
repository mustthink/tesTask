package db

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

func OpenDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(5)

	return db, nil
}

var ErrNoRecord = errors.New("error: подходящей записи не найдено")
