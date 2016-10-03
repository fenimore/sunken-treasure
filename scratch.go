package main

import (
	"fmt"

	"github.com/polypmer/sunken/api"
	"github.com/polypmer/sunken/database"
)

//"github.com/polypmer/sunken/geo"

func main() {
	//coordinates, _ := geo.Resolve("10017")
	//fmt.Println(coordinates[0], coordinates[1])

	//addr, _ := geo.Reverse(coordinates)
	//fmt.Println(addr)
	fmt.Println("Server takes a second to Load... Why?")

	// InitDB opens chest.db
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
