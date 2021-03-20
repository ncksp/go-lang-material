package main

import (
	"fmt"
)

type Student struct {
	name, address string
	age           int
}

func (student Student) toString() {
	fmt.Printf("My name is %s", student.name)
}

func main() {
	var nicko Student

	nicko.name = "Nicko"
	nicko.address = "indonesia"
	nicko.age = 1

	fmt.Print(nicko)

	student1 := Student{
		name:    "Student 1",
		address: "world",
		age:     10,
	}
	fmt.Print(student1)

	student2 := Student{"Student 2", "Indo", 10}
	fmt.Print(student2)
	student1.toString()
}
