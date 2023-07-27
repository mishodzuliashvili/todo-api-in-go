package main

import (
	"log"
	"net/http"
)

func main() {

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", Index)
	// router.HandleFunc("/todos/", TodoIndex)
	// router.HandleFunc("/todos/{todoId}", TodoSingle)

	// log.Fatal(http.ListenAndServe(":8080", router))
	// fmt.Println("Listening...")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
