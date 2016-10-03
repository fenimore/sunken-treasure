package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/polypmer/sunken/database"
)

// Global Database Connection Variable.
// TODO: Move to Context?
// TODO: Move to DBCon in database package?
var db *sql.DB

func Serve(connection *sql.DB) {
	db = connection
	fmt.Println("Serving API on port 8080")
	// A gorilla mux server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/stuff/{id}", ShowStuff)
	router.HandleFunc("/new", NewStuff)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index Page")
}

func ShowStuff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // a map of parameters
	id := vars["id"]
	fmt.Fprintln(w, "Showing:", id)
}

func NewStuff(w http.ResponseWriter, r *http.Request) {
	err := database.NewStuff(db, "Couch", "22203")
	if err != nil {
		fmt.Fprintf(w, "Error New Stuff %s", err)
	}
	fmt.Fprintln(w, "New stuff Posted")
}
