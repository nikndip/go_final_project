package db

import (
	"database/sql"
	"errors"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT '',
    title VARCHAR NOT NULL DEFAULT '',
    comment TEXT NOT NULL DEFAULT '',
    repeat VARCHAR(128) NOT NULL DEFAULT ''
        CHECK (length(repeat) <= 128)
);

CREATE INDEX scheduler_date_idx ON scheduler (date);
`

var db *sql.DB

func Init(dbFile string) error {
	_, err := os.Stat(dbFile)
	install := errors.Is(err, os.ErrNotExist)

	if err != nil && !install {
		return err
	}

	database, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	if err := database.Ping(); err != nil {
		database.Close()
		return err
	}

	if install {
		if _, err := database.Exec(schema); err != nil {
			database.Close()
			return err
		}
	}

	db = database
	return nil
}
