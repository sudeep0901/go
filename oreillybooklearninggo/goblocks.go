package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 10;
	if x > 5 {
		fmt.Println(x);
		x:= 5;
		fmt.Println(x);
	}
	fmt.Println(x);

	if n:= rand.Intn(10); n==0 {
		fmt.Println("That's too low",n);
	} else {
		fmt.Println("printing n:", n);
	}

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i);
	}

	i := 1;
	for i < 100 {
		fmt.Print(i)
		i = i * 2
	}

	evenVals := []int{2, 4, 6, 8, 10, 12}

	for i, v := range evenVals {
		fmt.Println(i, v);
	}


	a := 10;
	b := 20;

	fmt.Println(add(a , b));

	fmt.Println(a, b);

	variadic("Header", 1, 2, 3, 4, 5);
	// Infinite for loop
	// for {
	// 	fmt.Print("Hellow");
	// }
}
func add(a int, b int) int {
	// a = 100;
	return a + b;
}

// Go supports variadic parameters

func variadic(header string, vals ...int)  {
	val := make([]int, 0,len(vals))
	fmt.Println(val);
	for _ ,v := range val {
		fmt.Println(v * v);
	}
}