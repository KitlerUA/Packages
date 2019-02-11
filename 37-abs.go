package main

import (
	"fmt"
	"path/filepath"
)

//START OMIT

func main() {
	paths := []string{
		"/home/vkit/go/src/",
		"/b/c",
		"knok-kno-knok-penny",
	}

	fmt.Println("On Unix:")
	for _, p := range paths {
		rel, err := filepath.Abs(p)
		fmt.Printf("%q: %q %v\n", p, rel, err)
	}

}

//END OMIT
