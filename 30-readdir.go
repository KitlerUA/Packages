package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

//START OMIT

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

//END OMIT
