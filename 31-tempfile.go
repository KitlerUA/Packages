package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//START OMIT

func main() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	fmt.Println(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

//END OMIT
