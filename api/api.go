package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/polypmer/sunken/database"
)

func Serve() {
	fmt.Println("Serving API on port 8080")
	// A gorilla mux server
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/stuff/{id}", ShowStuff)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB()
	if err != nil {
		fmt.Fprintf(w, "Error with database init %s\n", err)
	}
	err = database.CreateTable(db)
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
