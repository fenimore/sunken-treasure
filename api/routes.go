package api

import "net/http"
import "github.com/gorilla/mux"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// NewRouter returns a mux.Router with
// the neatly defined routes defined etc.
func NewRouter() *mux.Router {
	// Why strict slash?
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}

	return router
}

// Handlers are in... handlers.go
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ShowStuff",
		"GET",
		"/stuff/{id}",
		ShowStuff,
	},
	Route{
		"NewStuff",
		"GET", // Change to post?
		"/new",
		NewStuff,
	},
	Route{
		"StuffIndex",
		"GET",
		"/all",
		StuffIndex,
	},
}
