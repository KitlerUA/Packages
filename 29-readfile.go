package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

//START OMIT

func main() {
	ioutil.WriteFile("knok-kno-knok-penny/hello", []byte("sksksksksksf"), 0755)

	content, err := ioutil.ReadFile("knok-kno-knok-penny/hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)

}

//END OMIT
