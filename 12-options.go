package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

//START OMIT
func main() {
	response, err := http.PostForm("http://localhost:8787/upload", url.Values{"qwerty": []string{"bembem"}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", response)

	client := http.Client{
		Timeout: 10,
	}
	response, err = client.PostForm("http://localhost:8787/upload", url.Values{"qwerty": []string{"bembem"}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", response)
}

//END OMIT
