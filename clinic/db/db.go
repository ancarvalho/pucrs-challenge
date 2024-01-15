package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var usersTable = `

CREATE TABLE IF NOT EXISTS users (
	id TEXT PRIMARY KEY,
	name VARCHAR(256),
	phone int
)

`

var appointmentsTable = `

CREATE TABLE IF NOT EXISTS appointments (
	id TEXT PRIMARY KEY,
	time TIMESTAMP,
	user_id TEXT,
	FOREIGN KEY (user_id) REFERENCES users(id)
)

`

func InstanciateDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", "./persistent.db")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = db.Exec(usersTable)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(appointmentsTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}
