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
		"./b/c",
	}
	base := "/home/vkit/go/src/github.com/KitlerUA/Packages"

	fmt.Println("On Unix:")
	for _, p := range paths {
		rel, err := filepath.Rel(base, p)
		fmt.Printf("%q: %q %v\n", p, rel, err)
	}

}

//END OMIT
