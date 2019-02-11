package main

import (
	"net/http"
)

//START OMIT
func main() {
	http.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)

			w.Write([]byte("u shel not pes"))
			return
		}

		w.Header().Add("Request-ID", "12345678")
		w.Write("cam in")

	})

	http.ListenAndServe("localhost:6777", nil)
}

//END OMIT
