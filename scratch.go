package main

import (
	"fmt"

	"github.com/polypmer/sunken/api"
	"github.com/polypmer/sunken/database"
)

//"github.com/polypmer/sunken/geo"

func main() {
	db, err := database.InitDB()
	if err != nil {
		fmt.Printf("Error with database init %s\n", err)
	}
	err = database.CreateTable(db)
	if err != nil {
		fmt.Printf("Error with database creation %s\n", err)
	}
	api.Serve(db)
}
