package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type myStruct struct{}

//START OMIT
func main() {
	client := http.Client{
		Timeout: 10,
	}
	response, err := client.Get("http://localhost:7047/currency")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

//END OMIT
