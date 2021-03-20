package main

import "fmt"

type Student struct {
	name, address string
	age           int
}

type Attribute interface {
	GetAge() int
}

func SetString(s string) interface{} {
	return s
}

func PrintAge(attribute Attribute) {
	fmt.Print("Umur ", attribute.GetAge())
}

func (s Student) GetAge() int {
	return s.age
}

func main() {
	student1 := Student{
		name:    "Student 1",
		address: "world",
		age:     10,
	}

	PrintAge(student1)

	s := SetString("nicko")

	fmt.Print(s)
}
