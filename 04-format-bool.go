package main

import (
	"fmt"
	"strconv"
)

//START OMIT
func main() {
	str := strconv.FormatBool(1 == 1)
	fmt.Println(str)

	str = strconv.FormatBool(false)
	fmt.Println(str)
}

//END OMIT
