package main

import (
	"fmt"
)

func main() {

	var x float64
	var y float64

	x = 1
	y = 2

	z := 100.9
	//declare but not used  is compile error
	n := 0
	fmt.Printf("x=%v, tyoe of %T\n", x, x)
	fmt.Printf("x=%v, type of %T\n", y, y)
	fmt.Printf("x=%v", z)
	var mean float64

	mean = (x + y) / 2.0
	fmt.Printf("resut: %v, type of %T\n", mean, mean)

}
