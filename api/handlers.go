package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/polypmer/sunken/database"
)

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
