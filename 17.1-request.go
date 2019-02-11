package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//START OMIT
func main() {
	http.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			return
		}
		r.ParseForm()

		mtime := r.Form.Get("mtime")

		m, err := strconv.Atoi(mtime)
		fmt.Println(m, err)

		fmt.Println(r.FormValue("mtime"))

	})

	http.ListenAndServe("localhost:6777", nil)
}

//END OMIT
