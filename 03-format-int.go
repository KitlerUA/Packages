package main

import (
	"fmt"
	"strconv"
)

//START OMIT
func main() {
	str := strconv.FormatInt(47, 10)
	fmt.Println(str)

	str = strconv.FormatInt(493, 8)
	fmt.Println(str)
}

//END OMIT
