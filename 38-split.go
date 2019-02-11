package main

import (
	"fmt"
	"path/filepath"
)

//START OMIT

func main() {
	paths := []string{
		"/home/arnie/amelia.jpg",
		"/mnt/photos/",
		"rabbit.jpg",
		"/usr/local//go",
		"/usr/local//go:/home",
	}
	fmt.Println("On Unix:")
	for _, p := range paths {
		dir, file := filepath.Split(p)
		fmt.Printf("input: %q\n\tdir: %q\n\tfile: %q\n", p, dir, file)

		parts := filepath.SplitList(p)
		fmt.Println("List:", parts)
	}
}

//END OMIT
