package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num = 150
	str := strconv.Itoa(num)
	fmt.Printf("Type : %T\nValue : %v\n", str, str)
}
