package main

import (
	"fmt"
)

func main() {
	fmt.Println("Frame memory allocated from go and varialbes defined into it")	
	c := 10
	d := 20

	// Passing addrss of varialbe
	fmt.Println(c, &c)
	increment(&c, &d)
	fmt.Println(c, &c)

	fmt.Println(d, &d)
	increment(&c, &d)
	fmt.Println(d, &d)

	fmt.Println(c, &c, d, &d)

}

func increment(rc *int, rd *int) {
	*rc++

	fmt.Println("Reference pointer")
	fmt.Println(rc)
	fmt.Println(rd)
	address(rc, rd)

	fmt.Println("After calling address function")
	fmt.Println(rc,*rc, rd,*rd)
	*rc++

}
// aDdress are also pass by value, causes mutating and voilets immuteability
func address(rc *int, rd *int) {
	rc = rd
	fmt.Println("In the address function")

	fmt.Println(rc,*rc, rd,*rd)
	fmt.Println(rd)

	fmt.Println(*rc, *rd)
	*rc++
	fmt.Println(*rc, *rd)

	*rd++
	fmt.Println(*rc, *rd)
	*rc++
	fmt.Println(*rc, *rd)

	*rd++
	fmt.Println(*rc, *rd)

}