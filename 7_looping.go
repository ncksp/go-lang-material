package main

import (
	"fmt"
)

func main() {
	count := 1

	for count <= 10 {
		count++
	}
	fmt.Print(count)

	for i := 0; i < count; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i)
	}

	names := []string{"nicko", "putra"}

	for _, name := range names {
		fmt.Print(name)
	}
}
