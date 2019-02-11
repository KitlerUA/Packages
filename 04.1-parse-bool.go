package main

import (
	"fmt"
	"strconv"
)

//START OMIT
func main() {
	str, err := strconv.ParseBool("1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)

	str, err = strconv.ParseBool("False")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}

//END OMIT
