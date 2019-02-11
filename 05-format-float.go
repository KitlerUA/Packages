package main

import (
	"fmt"
	"strconv"
)

//START OMIT
func main() {
	str := strconv.FormatFloat(1234.5667, 'f', 2, 64)
	fmt.Println(str)

	str = strconv.FormatFloat(1234.5667, 'e', -1, 64)
	fmt.Println(str)
}

//END OMIT
