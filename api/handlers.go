package api

import (
	"encoding/json"
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
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		fmt.Fprintln(w, "Error JSON encoding Stuff: %s", err)
	}
}

func StuffIndex(w http.ResponseWriter, r *http.Request) {
	stuffs, err := database.ReadStuffs(db)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(stuffs)
	if err != nil {
		fmt.Fprintln(w, "Error JSON encoding Stuffs: %s", err)
	}
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
