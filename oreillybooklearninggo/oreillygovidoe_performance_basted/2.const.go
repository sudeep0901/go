package main

import (
	"fmt"
)

func main() {
	const pi = 3.14
	fmt.Println("Testing constants in go")

	fmt.Println(pi)

	// const(
	// 	A1 = iota
	// 	B1 = iota
	// 	C1 = iota
	// 	D1 = iota
		
	// )

	const(
		A1 = iota
		B1
		C1
		D1
		
	)
	fmt.Println(A1, B1, C1, D1)
	if A1 == 0 {
		fmt.Println(A1)
	}

	const(
		A2 = 1>>iota
		B2
		C2
		D2
		
	)
	fmt.Println(A2, B2, C2, D2)

}
