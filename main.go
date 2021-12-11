package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// Not mapped
	// austin := person{"Euijin", "Yang"}

	// Mapped
	austin := person{firstName: "Euijin", lastName: "Yang"}

	fmt.Println(austin)
}
