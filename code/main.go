package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	var i int
	i = 42
	fmt.Println(i)
	var f float32
	// f = 3.14
	firstName := "Sudeep"
	fmt.Println(firstName, f)
	b := true
	fmt.Println(b)

	// pointers
	var lastName *string = new(string)

	fmt.Println(lastName)
	*lastName = "Sudeep Patel pointer"
	fmt.Println(*lastName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)
	firstName = "Manasvi"
	fmt.Println(ptr, *ptr)

	// const
	const pi float32 = 3.1415
	fmt.Println(pi)
	// pi = 4.314

	// piptr := &pi
	// fmt.Println(piptr)
	fmt.Println(pi + 3)
	fmt.Println(pi)

}
