package main

import (
	"io"
	"log"
	"os"
	"strings"
)

//START OMIT

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}

//END OMIT
