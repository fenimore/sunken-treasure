package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {
	// A gorilla mux server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/create", CreateDb)
	router.HandleFunc("/stuff/{id}", ShowStuff)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index Page")
}

func CreateDb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Database created")
}

func ShowStuff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // a map of parameters
	id := vars["id"]
	fmt.Fprintln(w, "Showing:", id)
}
