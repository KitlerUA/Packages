package main

import (
	"fmt"
	"strconv"
)

//START OMIT
func main() {
	str := strconv.Itoa(322)
	fmt.Println(str)

	str = strconv.Itoa(322e3)
	fmt.Println(str)
}

//END OMIT
