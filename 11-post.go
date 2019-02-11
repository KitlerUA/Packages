package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

//START OMIT
func main() {
	b := make([]byte, 5)
	response, err := http.Post("http://localhost:8787/upload", "application/json", bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", response)

	client := http.Client{
		Timeout: 10,
	}
	response, err = client.Post("http://localhost:8787/upload", "application/json", bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", response)
}

//END OMIT
