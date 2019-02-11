package main

import (
	"fmt"
	"net/http"
)

//START OMIT
func main() {
	http.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {

		coo, err := r.Cookie("token")
		fmt.Println(coo.Expires)
		fmt.Println(coo.String())
		fmt.Println(coo.Name)

	})

	http.ListenAndServe("localhost:6777", nil)
}

//END OMIT
