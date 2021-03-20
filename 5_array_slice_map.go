package main

import (
	"fmt"
)

func main() {
	var names [3]string

	names[0] = "nicko"
	names[1] = "putra"
	names[2] = "42"

	fmt.Print(names)

	values := [3]int{
		90,
		10,
		20,
	}

	fmt.Print(values)

	numbers := [...]int{
		10,
		20,
	}
	fmt.Print(numbers)

	//slice
	slice := names[1:3]

	fmt.Print(slice[1])

	newSlice := make([]string, 10, 10)
	newSlice[9] = "halo"

	fmt.Print(newSlice)

	newSliceInit := []int{1, 2, 3, 4, 2}
	fmt.Print(newSliceInit)

	//map
	books := make(map[string]string)

	books["name"] = "jos"
	books["author"] = "IPPB"

	delete(books, "name")

	fmt.Print(books)
}
