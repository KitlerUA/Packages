package main

import (
	"fmt"
	"strconv"
)

//START OMIT
func main() {
	str := strconv.FormatUint(47, 10)
	fmt.Println(str)

	str = strconv.FormatUint(493, 8)
	fmt.Println(str)
}

//END OMIT
