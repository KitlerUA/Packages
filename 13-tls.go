package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//START OMIT
func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{
		Transport:     tr,
		CheckRedirect: func(r *http.Request, via []*http.Request) error { return nil },
		Jar:           nil}

	response, err := client.Get("https://localhost:8787/upload")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", response)
}

//END OMIT
