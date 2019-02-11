package main

import (
	"net/http"
)

//START OMIT
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// do something
	})

	http.ListenAndServe("localhost:5387", mux)

	http.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {
		// do something
	})

	http.ListenAndServe("localhost:6777", nil)
}

//END OMIT
