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
	router.HandleFunc("/all", StuffIndex)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}

// Index handlers api root
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index Page")
}

// ShowStuff displays a stuff by ID
func ShowStuff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // a map of parameters
	id := vars["id"]
	fmt.Fprintln(w, "Showing:", id)
}

func StuffIndex(w http.ResponseWriter, r *http.Request) {
	stuffs, err := database.ReadStuffs(db)
	if err != nil {
		fmt.Println(err)
	}
	for _, stuff := range stuffs {
		fmt.Println(stuff)
	}
	fmt.Fprintln(w, "Stuff Index")
}

// NewStuff creates a new stuff
// TODO: take in JSON data
func NewStuff(w http.ResponseWriter, r *http.Request) {
	err := database.NewStuff(db, "Table", "10017")
	if err != nil {
		fmt.Fprintf(w, "Error New Stuff %s", err)
	}
	fmt.Fprintln(w, "New stuff Posted")
}
