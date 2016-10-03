// chest package, as in treasure chest, is the database package.
package database

import (
	"database/sql"

	"github.com/polypmer/sunken/geo"

	_ "github.com/mattn/go-sqlite3"
)

/* Database Helpers */

// InitDb Open()s a sqlite3 in a path.
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./chest.db")
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
    lat FLOAT NOT NULL,
    lon FLOAT NOT NULL,
    contact TEXT NOT NULL,
    date DATETIME
);
`
	_, err := db.Exec(sql_table)
	if err != nil {
		return err
	}
	return nil
}

// NewStuff Creates a new stuff object
func NewStuff(db *sql.DB, title, zip string) error {
	coord, err := geo.Resolve(zip)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO stuffs(title, zip, lat, lon," +
		"date, contact)values(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	// TODO: Generate Date/time
	res, err := stmt.Exec(title, zip, coord[0],
		coord[1], "1989-01-01", "555-555-5555")
	if err != nil {
		return err
	}
	// TODO: Look up this method
	_, err = res.LastInsertId()
	// Returns id
	if err != nil {
		return err
	}
	return nil
}

// TODO:
// UPDATE
// READ
// DELETE
