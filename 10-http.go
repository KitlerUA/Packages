package main

import (
	"fmt"
	"log"
	"net/http"
)

//START OMIT
func main() {
	response, err := http.Get("http://localhost:8787/upload")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", response)

	client := http.Client{
		Timeout: 10,
	}
	response, err = client.Get("http://localhost:8787/upload")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", response)
}

//END OMIT
