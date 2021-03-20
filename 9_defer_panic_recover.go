package main

import (
	"fmt"
)

func log() {
	fmt.Print("write log")
}

func run(err bool) {
	defer endApp()
	defer log()
	fmt.Print("running")

	if err {
		panic("ERR!")
	}

}

func endApp() {
	fmt.Print("end")

	message := recover()

	fmt.Print("Error occured", message)
}

func main() {
	run(true)
}
