package main

import (
	"fmt"
)

func hallo() {
	fmt.Print("Halooo")
}

func introMe(name string, age int) {
	fmt.Printf("saya %s umur saya %d", name, age)
}

func getMe(name string) string {
	return "nama saya " + name
}

func getMeMore() (string, string) {
	return "nicko", "putra"
}

func sum(numbers ...int) int {
	count := 0
	for _, num := range numbers {
		count += num
	}
	return count
}

func main() {
	hallo()
	introMe("nicko", 1)
	fmt.Print(getMe("nicko"))

	fmt.Print(getMeMore())

	fmt.Print(sum(1, 23, 123, 123, 1, 2))

	sumNumber := sum
	fmt.Print(sumNumber(12, 12, 312, 31, 23, 12))
}
