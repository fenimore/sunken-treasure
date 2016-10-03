package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/polypmer/database"
)

func Serve() {
	// A gorilla mux server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/stuff/{id}", ShowStuff)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB()
	if err != nil {
		fmt.Fprintf(w, "Error with database init %s\n", err)
	}
	err := database.CreateTable()
	if err != nil {
		fmt.Fprintf(w, "Error with database creation %s\n", err)
	}
	fmt.Fprintln(w, "Index Page:\nDatabase Created if not already.")
}

func ShowStuff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // a map of parameters
	id := vars["id"]
	fmt.Fprintln(w, "Showing:", id)
}
