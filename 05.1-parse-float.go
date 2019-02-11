package main

import (
	"fmt"
	"strconv"
)

//START OMIT
func main() {
	str, err := strconv.ParseFloat("1", 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)

	str, err = strconv.ParseFloat("False", 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}

//END OMIT
