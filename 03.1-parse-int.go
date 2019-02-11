package main

import (
	"fmt"
	"strconv"
)

//START OMIT
func main() {
	str, err := strconv.ParseInt("755", 8, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)

	str, err = strconv.ParseInt("A", 16, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}

//END OMIT
