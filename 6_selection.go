package main

import (
	"fmt"
)

func main() {
	name := "putra"

	if name == "putra" {
		fmt.Print("bener")
	} else {
		fmt.Print("salah")
	}

	switch name {
	case "nicko":
		fmt.Print("bener")
	default:
		fmt.Print("salah")
	}
}
