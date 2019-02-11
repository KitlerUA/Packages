package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

//START OMIT
func main() {
	u, _ := url.Parse("http://localhost:8787/upload")
	r := http.Request{
		Method: http.MethodGet,
		URL:    u,
		Header: http.Header{"Range": []string{"0-9/10"}},
	}

	client := http.Client{Timeout: 10}

	response, err := client.Do(&r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", response)
}

//END OMIT
