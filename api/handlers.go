package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/polypmer/sunken/database"
)

// Index handlers api root
func Index(w http.ResponseWriter, r *http.Request) {
	// Go's net/http package tries to guess output
	// type, but cause I know it (It'll be json)
	// I'll set it myself.
	fmt.Fprintln(w, "Index Page")
}

// ShowStuff displays a stuff by ID
func ShowStuff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // a map of parameters
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	s, err := database.ReadStuff(db, id)
	if err != nil {
		fmt.Fprintln(w, "Error Reading Stuff: %s", err)
	}
	fmt.Fprintf(w, "Showing: %s at %s\n", s.Title, s.Zip)
}

func StuffIndex(w http.ResponseWriter, r *http.Request) {
	stuffs, err := database.ReadStuffs(db)
	if err != nil {
		fmt.Println(err)
	}
	for _, stuff := range stuffs {
		fmt.Fprintf(w, "%d Stuff: %s\n  Zip: %s\n", stuff.Id, stuff.Title, stuff.Zip)
	}
	//fmt.Fprintln(w, "Stuff Index")
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
