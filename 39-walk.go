package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//START OMIT

func main() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			fmt.Printf("Directory: %s\n", info.Name())
			return nil
		}
		fmt.Printf("File %s with size %d\n", info.Name(), info.Size())
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path")
		return
	}
}

//END OMIT
