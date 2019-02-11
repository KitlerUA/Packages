package main

import (
	"fmt"
	"net/http"
)

//START OMIT
func main() {
	http.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		r.ParseMultipartForm()

		file, err := r.MultipartForm.File["activity_log.csv"][0].Open()

		dst := make([]byte, 1024)

		_, err = file.Read(dst)

		fmt.Println(r.MultipartForm.Value["mtime"][0])

	})

	http.ListenAndServe("localhost:6777", nil)
}

//END OMIT
