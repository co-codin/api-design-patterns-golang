package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	maxOpenDbConn = 25
	maxIdleDBconn
	maxDBLifetime = 5 * time.Minute
)

func initMySQLDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDBconn)
	db.SetConnMaxLifetime(maxDBLifetime)

	return db, nil
}
