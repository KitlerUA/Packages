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
	s := io.NewSectionReader(r, 5, 17)

	if _, err := io.Copy(os.Stdout, s); err != nil {
		log.Fatal(err)
	}

}

//END OMIT
