package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/polypmer/sunken/database"
	"github.com/polypmer/sunken/geo"
	"github.com/polypmer/sunken/stuff"
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
	var stuff stuff.Stuff
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // wtf num?
	// That number protects agains huge json
	if err != nil {
		fmt.Println(err)
	}
	err = r.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	// Unmarshal means take the body json thingy
	// and stick it into one of my fancy structs
	err = json.Unmarshal(body, &stuff)
	if err != nil {
		// IF shit doesn't work out
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable data
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			fmt.Println(err)
		}
	}

	coord, err := geo.Resolve(stuff.Zip)
	if err != nil {
		fmt.Println(err)
	}
	stuff.Lat, stuff.Lon = coord[0], coord[1]
	stuff.Date = time.Now()

	id, err := database.NewStuff(db, stuff)
	if err != nil {
		fmt.Println(err)
	}
	stuff.Id = id
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(stuff)
	// I send back the created Stuff in json so
	// that way the client has access to the ID I created
	if err != nil {
		fmt.Println(err)
	}
}
