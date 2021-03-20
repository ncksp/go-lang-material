package main

import (
	"fmt"
)

func main() {
	a, b := 10, 15
	c := a + b

	fmt.Print(c)

	c += 1
	fmt.Print(c)

	c++
	fmt.Print(c)

	compare := a == b
	fmt.Print(compare)

}
