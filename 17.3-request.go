package main

import (
	"fmt"
	"net/http"
)

//START OMIT
func main() {
	http.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Header.Get("Host"))

		fmt.Println(r.Host)
		fmt.Println(r.URL)
		fmt.Println(r.RemoteAddr)

		user, password := r.BasicAuth()
		fmt.Println(user, password)

	})

	http.ListenAndServe("localhost:6777", nil)
}

//END OMIT
