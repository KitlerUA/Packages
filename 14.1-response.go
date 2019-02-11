package main

import (
	"fmt"
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

	allMyCookies := response.Cookies()
	for i := range allMyCookies {
		fmt.Println(allMyCookies[i].String(), allMyCookies[i].Expires)
	}
}

//END OMIT
