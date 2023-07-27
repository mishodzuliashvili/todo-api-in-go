package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := Logger(route.HandlerFunc, route.Name)
		// i could pass handler itself but i need logger
		router.
			Methods(route.HTTPMethod).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
