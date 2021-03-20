package main

import (
	"fmt"
)

func main() {
	var num int32 = 123
	var boolean bool = true
	var s string = "nicko"

	va1, var2 := "hehe", "haha"

	const var3 string = "hihi"

	const (
		var4 string = "hoho"
		var5 uint32 = 1000
	)

	type nohp string

	var no nohp = "123444"

	fmt.Print(num)
	fmt.Print(boolean)
	fmt.Print(s, va1, var2, var3, var4, var5, no)
}
