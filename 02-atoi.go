package main

import (
	"fmt"
	"strconv"
)

//START OMIT
func main() {
	str, err := strconv.Atoi("0322")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)

	str, err = strconv.Atoi("322f")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}

//END OMIT
