// chest package, as in treasure chest, is the database package.
package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/polypmer/sunken/stuff"
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
    zip TEXT,
    lat FLOAT NOT NULL,
    lon FLOAT NOT NULL,
    contact TEXT NOT NULL,
    date DATETIME,
    expired BOOLEAN DEFAULT FALSE
);
`
	_, err := db.Exec(sql_table)
	if err != nil {
		return err
	}
	return nil
}

// NewStuff Creates a new stuff object
func NewStuff(db *sql.DB, stuff stuff.Stuff) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO stuffs(title, zip, lat," +
		"lon, date, contact)values(?,?,?,?,?,?)")
	if err != nil {
		return -1, err
	}
	// TODO: Generate Date/time
	res, err := stmt.Exec(stuff.Title, stuff.Zip, stuff.Lat,
		stuff.Lon, stuff.Date, stuff.Contact)
	if err != nil {
		return -1, err
	}
	// TODO: Look up this method
	// TODO: change to random hash
	id, err := res.LastInsertId()
	// Returns id
	if err != nil {
		return -1, err
	}
	return id, nil
}

// TODO:
// UPDATE
// READ
// ReadStuff returns a stuff by id.
func ReadStuff(db *sql.DB, id int) (stuff.Stuff, error) {
	rows, err := db.Query("select * from stuffs where id = ?", id)
	s := stuff.Stuff{}
	if err != nil {
		return s, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&s.Id, &s.Title,
			&s.Zip, &s.Lat, &s.Lon,
			&s.Contact, &s.Date,
			&s.Expired)
	}
	rows.Close()
	return s, nil
}

// ReadStuffs returns a slice all stuffs.
// TODO: make only for active stuffs.
func ReadStuffs(db *sql.DB) ([]stuff.Stuff, error) {
	stuffs := make([]stuff.Stuff, 0)

	stmt := "SELECT * FROM stuffs"
	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := stuff.Stuff{}
		err = rows.Scan(&s.Id, &s.Title,
			&s.Zip, &s.Lat, &s.Lon,
			&s.Contact, &s.Date,
			&s.Expired)
		if err != nil {
			return nil, err
		}
		stuffs = append(stuffs, s)
	}
	rows.Close() // Redundant but good
	return stuffs, nil
}

// DELETE
