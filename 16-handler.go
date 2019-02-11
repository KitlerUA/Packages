package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

//START OMIT
func main() {
	mux := http.NewServeMux()

	mux.Handle("/user", HandlerWithMiddleware{})

	http.ListenAndServe(":2323", mux)
}

type HandlerWithMiddleware struct{}

func (h HandlerWithMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, "test:", 0)
	start := time.Now()
	func(w ResponseWriter, r *Request) {
		if r.Method != http.MethodGet {
			return
		}
		// do something else
	}(w, r)
	end := time.Now()
	logger.Println(end.Sub(start))
}

//END OMIT
