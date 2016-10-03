package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Global Database Connection Variable.
// TODO: Move to Context?
// TODO: Move to DBCon in database package?
var db *sql.DB

// Serve Requires a DB connection to be passed in.
// The database is a global pool, protecting against
// concurrent writes?
func Serve(connection *sql.DB) {
	db = connection
	fmt.Println("Serving API on port 8080")
	// A gorilla mux server
	router := NewRouter() // check routes.go
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
