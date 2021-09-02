package main

import (
	"fmt"
)

type Employee struct {
	firstname string
	lastname  string
	city      string
	pincode   int
	permanent bool
}

func main() {
	fmt.Println("Process started")

	emp1 := Employee{"sudeep", "patel", "Ratlam", 457001, true}
	fmt.Println(emp1)
	fmt.Println(emp1.firstname, emp1.lastname, emp1.city, emp1.pincode)

	emp2 := Employee{}
	fmt.Println(emp2)
	fmt.Println(emp2.firstname)

}
