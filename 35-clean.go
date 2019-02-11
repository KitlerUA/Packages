package main

import (
	"fmt"
	"path/filepath"
)

//START OMIT

func main() {
	fmt.Println(filepath.Clean("/a/b/../b/.."))
	fmt.Println(filepath.Clean("a"))
	fmt.Println(filepath.Clean(""))
	fmt.Println(filepath.Clean("a//b/"))
}

//END OMIT
