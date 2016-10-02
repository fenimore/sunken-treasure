package main

import (
	"fmt"

	"github.com/polypmer/sunken/geo"
)

func main() {
	coordinates, _ := geo.Resolve("10017")
	fmt.Println(coordinates[0], coordinates[1])
}
