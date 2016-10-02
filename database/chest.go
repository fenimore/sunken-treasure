// chest package, as in treasure chest, is the database package.
package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/* Database Helpers */

// InitDb Open()s a sqlite3 in path.
func InitDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateTable in the sql.DB
func CreateTable(db *sql.DB) error {
	// TODO: add Binary Pictures
	// TODO: add coordinate floats
	// TODO: is zip int or string?
	// TODO: Date and time
	sql_table := `
CREATE TABLE IF NOT EXISTS stuffs(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,                
    zip TEXT NOT NULL,
    contact TEXT NOT NULL,
    date DATE
);
`
	_, err := db.Exec(sql_table)
	if err != nil {
		return err
	}
}

// NewStuff Creates a new stuff object
func NewStuff(db *sql.DB, title, zip string) error {
	stmt, err := db.Prepare("INSERT INTO stuffs(title, zip," +
		"date)values(?,?,?)")
	if err != nil {
		return err
	}
	// TODO: Generate Date/time
	res, err := stmt.Exec(title, zip, "1989-01-01")
	if err != nil {
		return err
	}
	// TODO: Look up this method
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
}

// TODO:
// UPDATE
// READ
// DELETE
